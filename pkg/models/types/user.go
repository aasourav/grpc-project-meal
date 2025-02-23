package types

import "time"

type User struct {
	ID string `json:"id" bson:"_id,omitempty"`
	UserLogin
	Name         string     `json:"name" bson:"name" validate:"required,min=2"`
	EmployeeId   string     `json:"employeeId" bson:"employeeId" validate:"required"`
	WeeklyPlan   []bool     `json:"weeklyplan" bson:"weeklyPlan" validate:"required,len=7"` // last weekly plan for check every element
	CreatedAt    time.Time  `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time  `bson:"updatedAt" json:"updatedAt"`
	Department   Department `json:"department" bson:"department" validate:"required,oneof=TECHNOLOGY HR MARKETING FINANCE"`
	IsApproved   bool       `bson:"isApproved" json:"isApproved"`
	ApprovedById string     `bson:"approvedById" json:"approvedById"`
}

type UserLogin struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"-" bson:"password" validate:"required,min=6"`
}
