package interfaces

import (
	models "aas.dev/pkg/models/admin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminRepository interface {
	CreateAdmin(user *models.Admin) error
	DeleteAdminById(id string) error
	GetAdminByEmail(email string) (*models.Admin, error)
	UpdateAdminById(admin *models.Admin) error
	GetAdmins() (*[]models.Admin, error)
	GetAdminById(id primitive.ObjectID) (*models.Admin, error)
}
