package interfaces

import models "aas.dev/pkg/models/admin"

type AdminRepository interface {
	CreateAdmin(user *models.Admin) error
	DeleteAdminById(id string) error
	GetAdminByEmail(email string) (*models.Admin, error)
}
