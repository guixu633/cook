package repository

import (
	"server/internal/model"

	"gorm.io/gorm"
)

type MealRepository interface {
	Create(meal *model.Meal) error
	Update(meal *model.Meal) error
	Delete(id uint) error
	List() ([]model.Meal, error)
	FindByID(id uint) (*model.Meal, error)
}

type mealRepository struct {
	db *gorm.DB
}

func NewMealRepository(db *gorm.DB) MealRepository {
	return &mealRepository{db: db}
}

func (r *mealRepository) Create(meal *model.Meal) error {
	return r.db.Create(meal).Error
}

func (r *mealRepository) Update(meal *model.Meal) error {
	return r.db.Save(meal).Error
}

func (r *mealRepository) Delete(id uint) error {
	return r.db.Delete(&model.Meal{}, id).Error
}

func (r *mealRepository) List() ([]model.Meal, error) {
	var meals []model.Meal
	// Order by date descending by default
	err := r.db.Order("date desc, created_at desc").Find(&meals).Error
	return meals, err
}

func (r *mealRepository) FindByID(id uint) (*model.Meal, error) {
	var meal model.Meal
	err := r.db.First(&meal, id).Error
	if err != nil {
		return nil, err
	}
	return &meal, nil
}
