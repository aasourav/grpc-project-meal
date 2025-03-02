package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Verification struct {
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" validate:"required"`
}

var verification = validator.New()

func (a *Verification) VerifiactionValidate() error {
	return verification.Struct(a)
}
