package interfaces

import models "aas.dev/pkg/models/user"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}
