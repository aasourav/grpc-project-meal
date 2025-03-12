package models

import (
	"time"

	"aas.dev/pkg/models/types"
)

type Admin struct {
	ID                      string           `json:"id" bson:"_id,omitempty"`
	Name                    string           `json:"name" bson:"name" validate:"required"`
	Role                    types.AdminRole  `json:"role" bson:"role" validate:"required,oneof=SADMIN ADMIN"`
	DepartmentInCharge      types.Department `json:"departmentInCharge" bson:"departmentInCharge" validate:"required,oneof=TECHNOLOGY HR MARKETING FINANCE"`
	CreatedAt               time.Time        `bson:"createdAt" json:"createdAt"`
	UpdatedAt               time.Time        `bson:"updatedAt" json:"updatedAt"`
	EmployeeId              string           `json:"employeeId" bson:"employeeId" validate:"required"`
	PendingUserApprovalIds  []string         `bson:"pendingApprovalIds" json:"pendingApprovalIds"`
	IsEmailVerified         bool             `bson:"isEmailVerified" json:"isEmailVerified"`
	IsApproved              bool             `bson:"isApproved" json:"isApproved"`
	PendingAdminApprovalIds []string         `bson:"pendingAdminApprovalIds" json:"pendingAdminApprovalIds"`
	Department              types.Department `json:"department" bson:"department" validate:"required,oneof=TECHNOLOGY HR MARKETING FINANCE"`
	Email                   string           `json:"email" bson:"email" validate:"required,email"`
	Password                string           `json:"password" bson:"password" validate:"required,min=6"`
}

type AdminLogin struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=6"`
}
