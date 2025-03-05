package models

import (
	"time"

	"aas.dev/pkg/models/types"
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID                   string                        `json:"id" bson:"_id,omitempty"`
	Name                 string                        `json:"name" bson:"name" validate:"required,min=2"`
	IsEmailApproved      bool                          `bson:"isEmailApproved" json:"isEmailApproved"`
	Email                string                        `json:"email" bson:"email" validate:"required,email"`
	Password             string                        `json:"password" bson:"password" validate:"required,min=6"`
	EmployeeId           string                        `json:"employeeId" bson:"employeeId" validate:"required"`
	WeeklyPlan           []bool                        `json:"weeklyplan" bson:"weeklyPlan" validate:"required,len=7"` // last weekly plan for check every element
	CreatedAt            time.Time                     `bson:"createdAt" json:"createdAt"`
	UpdatedAt            time.Time                     `bson:"updatedAt" json:"updatedAt"`
	Department           types.Department              `json:"department" bson:"department" validate:"required,oneof=TECHNOLOGY HR MARKETING FINANCE"`
	IsApproved           bool                          `bson:"isApproved" json:"isApproved"`
	ApprovedById         string                        `bson:"approvedById" json:"approvedById"`
	RequestNewWeeklyPlan *[]types.RequestNewWeeklyPlan `bson:"requestNewWeeklyPlan" json:"requestNewWeeklyPlan"`
}

type UserLogin struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=6"`
}

var userValidate = validator.New()

func (u *User) UserValidate() error {
	return userValidate.Struct(u)
}

/**
package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// User struct with optional WeeklyPlan validation
type User struct {
	ID         string  `json:"id" bson:"_id,omitempty"`
	Name       string  `json:"name" bson:"name" validate:"required,min=2"`
	Email      string  `json:"email" bson:"email" validate:"required,email"`
	Password   string  `json:"password" bson:"password" validate:"required,min=6"`
	EmployeeId string  `json:"employeeId" bson:"employeeId" validate:"required"`
	WeeklyPlan *[]bool `json:"weeklyPlan,omitempty" bson:"weeklyPlan,omitempty" validate:"omitempty,weeklyplan"` // Make it a pointer to allow nil
}

// Custom validation function for WeeklyPlan
func validateWeeklyPlan(fl validator.FieldLevel) bool {
	weeklyPlan, ok := fl.Field().Interface().(*[]bool)
	if !ok || weeklyPlan == nil {
		return true // ✅ Skip validation if nil
	}

	// Ensure it has exactly 7 elements
	if len(*weeklyPlan) != 7 {
		return false
	}

	// Count the number of `true` values
	trueCount := 0
	for _, day := range *weeklyPlan {
		if day {
			trueCount++
		}
	}

	// Ensure at least 3 days are `true`
	return trueCount >= 3
}

func main() {
	validate := validator.New()

	// Register custom validator
	_ = validate.RegisterValidation("weeklyplan", validateWeeklyPlan)

	// Test Case 1: WeeklyPlan is nil (Should Pass ✅)
	user1 := User{
		Name:       "John Doe",
		Email:      "john@example.com",
		Password:   "password123",
		EmployeeId: "EMP123",
		WeeklyPlan: nil, // ✅ No validation error
	}

	// Test Case 2: WeeklyPlan has valid values (Should Pass ✅)
	user2 := User{
		Name:       "John Doe",
		Email:      "john@example.com",
		Password:   "password123",
		EmployeeId: "EMP123",
		WeeklyPlan: &[]bool{true, false, true, true, false, false, false}, // ✅ Valid (3 `true`)
	}

	// Test Case 3: WeeklyPlan has only 2 `true` values (Should Fail ❌)
	user3 := User{
		Name:       "John Doe",
		Email:      "john@example.com",
		Password:   "password123",
		EmployeeId: "EMP123",
		WeeklyPlan: &[]bool{true, false, false, false, true, false, false}, // ❌ Invalid (only 2 `true`)
	}

	users := []User{user1, user2, user3}
	for i, user := range users {
		fmt.Printf("\nTesting User %d:\n", i+1)
		err := validate.Struct(user)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("Field:", err.Field(), "| Error:", err.Tag(), "| Param:", err.Param())
			}
		} else {
			fmt.Println("✅ Validation Passed!")
		}
	}
}
**/
