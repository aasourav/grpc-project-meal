package interfaces

import "aas.dev/pkg/models/types"

type MealRepository interface {
	AddMeal(meal types.MealConsume) error
	GetMealByUser(meal types.MealConsume) (*types.MealConsume, error)
}
