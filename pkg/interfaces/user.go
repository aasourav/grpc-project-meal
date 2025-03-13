package interfaces

import (
	models "aas.dev/pkg/models/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	DeleteUserById(id string) error
	GetUserById(id primitive.ObjectID) (*models.User, error)
	UpdateUserById(user *models.User) error
	UpdatePasswordById(admin *models.User) error
}
