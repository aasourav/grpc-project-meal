package services

import (
	"errors"
	"fmt"
	"time"

	"aas.dev/pkg/interfaces"
	"aas.dev/pkg/models/types"
	models "aas.dev/pkg/models/user"
	verificationModels "aas.dev/pkg/models/verification"

	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepo         interfaces.UserRepository
	verificationRepo interfaces.VerifiactionRepository
}

func NewUserService(userRepo interfaces.UserRepository, verificationRepo interfaces.VerifiactionRepository) *UserService {
	return &UserService{userRepo: userRepo, verificationRepo: verificationRepo}
}

func (s *UserService) RegisterUser(c *gin.Context, user *models.User) error {
	userDoc, _ := s.FindUserByEmail(user.Email)
	if userDoc != nil {
		return errors.New("email already exist")
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Password, _ = utils.HashPassword(user.Password)
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	userDoc, err = s.FindUserByEmail(user.Email)
	if err != nil {
		return errors.New("user find error: " + err.Error())
	}

	expires := time.Now().Add(time.Second * 120).Unix()
	jwt, _ := utils.GenerateJWT(user.Email, "userId", expires)
	verifyLink := utils.GetBaseURL(c) + fmt.Sprintf("/users/verify?u=%s", jwt)

	emailVerifyData := types.EmailVerifyTypes{
		Email:           user.Email,
		VerificaionLink: verifyLink,
		Name:            user.Name,
	}

	_, err = NewGeneralService(nil).EmailVerify(c, emailVerifyData)
	if err != nil {
		userDoc, _ = s.FindUserByEmail(user.Email)
		s.userRepo.DeleteUserById(userDoc.ID)
		return err
	}

	mailData := &verificationModels.Verification{
		UserId:    userDoc.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	s.verificationRepo.CreateVerificationRepo(mailData)
	return nil
}

func (s *UserService) FindUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}
