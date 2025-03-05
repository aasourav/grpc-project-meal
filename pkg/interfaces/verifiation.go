package interfaces

import models "aas.dev/pkg/models/verification"

type VerifiactionRepository interface {
	CreateVerificationRepo(verificationData *models.Verification) error
	GetVerificationDocByUserId(userId string) (*models.Verification, error)
	DeleteVeruficationByUserId(userId string) error
}
