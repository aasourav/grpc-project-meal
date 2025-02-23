package services

import (
	"errors"
	"time"

	"aas.dev/pkg/interfaces"
	models "aas.dev/pkg/models/user"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	userDoc, _ := s.FindUserByEmail(user.Email)
	if userDoc != nil {
		return errors.New("email already exist")
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	return s.repo.CreateUser(user)
}

func (s *UserService) FindUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}
