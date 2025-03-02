package interfaces

import models "aas.dev/pkg/models/verification"

type VerifiactionRepository interface {
	CreateVerificationRepo(verificationData *models.Verification) error
}
