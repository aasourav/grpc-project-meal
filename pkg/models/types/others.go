package types

type EmailVerifyTypes struct {
	Email string `json:"email" bson:"email" validate:"required,email"`
	Name  string `json:"name" bson:"name" validate:"required"`
}
