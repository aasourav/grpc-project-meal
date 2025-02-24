package models

import (
	"aas.dev/pkg/models/types"
	"github.com/go-playground/validator/v10"
)

type Admin struct {
	types.Admin
}

type AdminLogin struct {
	types.AdminLogin
}

var adminValidate = validator.New()

func (a *Admin) AdminValidate() error {
	return adminValidate.Struct(a)
}

func (a *AdminLogin) AdminLoginValidate() error {
	return adminValidate.Struct(a)
}
