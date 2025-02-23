package models

import (
	"aas.dev/pkg/models/types"
	"github.com/go-playground/validator/v10"
)

type Admin struct {
	types.Admin
}

var adminValidate = validator.New()

func (a *Admin) AdminValidateConsumeValidate() error {
	return adminValidate.Struct(a)
}
