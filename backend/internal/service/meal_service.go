package service

import (
	"server/internal/model"
	"server/internal/repository"
	"time"
)

type MealService interface {
	CreateMeal(name, description string, date time.Time, imageURLs []string) (*model.Meal, error)
	DeleteMeal(id uint) error
	ListMeals() ([]model.Meal, error)
}

type mealService struct {
	repo repository.MealRepository
}

func NewMealService(repo repository.MealRepository) MealService {
	return &mealService{repo: repo}
}

func (s *mealService) CreateMeal(name, description string, date time.Time, imageURLs []string) (*model.Meal, error) {
	meal := &model.Meal{
		Name:        name,
		Description: description,
		Date:        date,
		ImageURLs:   imageURLs,
	}
	if err := s.repo.Create(meal); err != nil {
		return nil, err
	}
	return meal, nil
}

func (s *mealService) DeleteMeal(id uint) error {
	return s.repo.Delete(id)
}

func (s *mealService) ListMeals() ([]model.Meal, error) {
	return s.repo.List()
}
