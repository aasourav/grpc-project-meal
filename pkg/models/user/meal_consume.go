package models

import (
	"aas.dev/pkg/models/types"
	"github.com/go-playground/validator/v10"
)

type MealConsume struct {
	types.MealConsume
}

var mealConsumeValidate = validator.New()

func (u *User) MealConsumeValidate() error {
	return mealConsumeValidate.Struct(u)
}
