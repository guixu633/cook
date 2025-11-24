package service

import (
	"fmt"
	"server/internal/model"
	"server/internal/pkg/oss"
	"server/internal/repository"
	"time"
)

type MealService interface {
	// imageFilenames are the temp filenames in storage
	CreateMeal(name, description string, date time.Time, imageFilenames []string) (*model.Meal, error)
	DeleteMeal(id uint) error
	ListMeals() ([]model.Meal, error)
}

type mealService struct {
	repo      repository.MealRepository
	ossClient *oss.Client
}

func NewMealService(repo repository.MealRepository, ossClient *oss.Client) MealService {
	return &mealService{
		repo:      repo,
		ossClient: ossClient,
	}
}

func (s *mealService) CreateMeal(name, description string, date time.Time, imageFilenames []string) (*model.Meal, error) {
	// 1. Create Meal record first to get the ID
	meal := &model.Meal{
		Name:        name,
		Description: description,
		Date:        date,
		ImageURLs:   []string{}, // Initially empty
	}

	if err := s.repo.Create(meal); err != nil {
		return nil, err
	}

	// 2. Move images from tmp to meal directory and get public URLs
	var finalURLs []string
	for _, filename := range imageFilenames {
		url, err := s.ossClient.MoveFromTmpToMeal(filename, meal.ID)
		if err != nil {
			// In a real world scenario, we might want to rollback or log this error
			// For now we just log and skip this image
			fmt.Printf("Failed to process image %s: %v\n", filename, err)
			continue
		}
		finalURLs = append(finalURLs, url)
	}

	// 3. Update Meal record with final URLs
	meal.ImageURLs = finalURLs
	// We need to save the updated meal. Since Create() only inserts, we can use Save() or Update() via repo.
	// But our current repo interface only has Create, Delete, List, FindByID.
	// Let's reuse Create logic if the underlying DB method is Save (upsert) or add an Update method.
	// GORM's Create usually errors on duplicate key.
	// Let's use GORM's Save or Updates on the object.
	// We'll need to extend the Repository interface or just use a direct update here if repo exposed DB, but better to extend Repo.
	// For now, I'll assume we can add an Update method to repository.
	
	// Hack for now: Re-using Create won't work.
	// Let's assume we need to update the repo.
	// I will fix the repository first.
	
	return meal, s.repo.Update(meal)
}

func (s *mealService) DeleteMeal(id uint) error {
	return s.repo.Delete(id)
}

func (s *mealService) ListMeals() ([]model.Meal, error) {
	return s.repo.List()
}
