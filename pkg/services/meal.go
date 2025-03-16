package services

import (
	"aas.dev/pkg/interfaces"
	"aas.dev/pkg/models/types"
)

type MealService struct {
	mealRepo interfaces.MealRepository
}

func NewMealService(mealRepo interfaces.MealRepository) *MealService {
	return &MealService{mealRepo: mealRepo}
}

func (service *MealService) AddMeal(meal types.MealConsume) error {
	return service.mealRepo.AddMeal(meal)
}

func (service *MealService) GetMealByUser(meal types.MealConsume) (*types.MealConsume, error) {
	return service.mealRepo.GetMealByUser(meal)
}
