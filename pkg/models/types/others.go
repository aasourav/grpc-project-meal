package types

import "time"

type EmailVerifyTypes struct {
	Email           string `json:"email" bson:"email" validate:"required,email"`
	Name            string `json:"name" bson:"name" validate:"required"`
	VerificaionLink string `json:"verificationLink" bson:"verificationLink" validate:"required"`
}

type Verification struct {
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" validate:"required"`
}

type ResetPassword struct {
	Email          string `json:"email" bson:"email" validate:"required,email"`
	CurrentPawword string `json:"currentPassword" bson:"currentPassword" validate:"required"`
	NewPassword    string `json:"newPassword" bson:"newPassword" validate:"required"`
}
