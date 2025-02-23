package types

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Department string
type AdminRole string

const (
	// if this constant group is modified then. must have to update `Role` validator field
	SUPERADMIN AdminRole = "SADMIN"
	ADMIN      AdminRole = "ADMIN"
)

const (
	// if this constant group is modified then. must have to update `RespectiveDepartment` validator field
	TECHNOLOGY Department = "TECHNOLOGY"
	HR         Department = "HR"
	MARKETING  Department = "MARKETING"
	FINANCE    Department = "FINANCE"
)

type Admin struct {
	ID                      string     `json:"id" bson:"_id,omitempty"`
	UserId                  string     `json:"userId" bson:"userId" validate:"required"`
	Name                    string     `json:"name" bson:"name" validate:"required"`
	Role                    AdminRole  `json:"role" bson:"role" validate:"required,oneof=SADMIN ADMIN"`
	DepartmentInCharge      Department `json:"departmentInCharge" bson:"departmentInCharge" validate:"required,oneof=TECHNOLOGY HR MARKETING FINANCE"`
	CreatedAt               time.Time  `bson:"createdAt" json:"createdAt"`
	UpdatedAt               time.Time  `bson:"updatedAt" json:"updatedAt"`
	EmployeeId              string     `json:"employeeId" bson:"employeeId" validate:"required"`
	PendingUserApprovalIds  []string   `bson:"pendingApprovalIds" json:"pendingApprovalIds"`
	PendingAdminApprovalIds []string   `bson:"pendingAdminApprovalIds" json:"pendingAdminApprovalIds"`
	Department              Department `json:"department" bson:"department" validate:"required,oneof=TECHNOLOGY HR MARKETING FINANCE"`
	AdminLogin
}

type AdminLogin struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"-" bson:"password" validate:"required,min=6"`
}

var adminValidate = validator.New()

func (u *User) AdminValidateConsumeValidate() error {
	return adminValidate.Struct(u)
}
