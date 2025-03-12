package services

import (
	"errors"
	"fmt"
	"time"

	"aas.dev/pkg/interfaces"
	"aas.dev/pkg/models/types"
	models "aas.dev/pkg/models/user"
	verificationModels "aas.dev/pkg/models/verification"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func (s *UserService) VerifyUser(c *gin.Context) error {
	if c.Query("u") == "" {
		return errors.New("invalid request")
	}

	userId, err := utils.VerifyJWT(c.Query("u"), "userId")
	if err != nil {
		return err
	}

	verificationData, err := s.verificationRepo.GetVerificationDocByUserId(fmt.Sprintf("%v", userId))
	if err != nil {
		return err
	}

	if time.Since(verificationData.CreatedAt) > types.VERIFICATION_EXPIRY_SECONDS*time.Second {
		return errors.New("verification link expired")
	}

	err = s.verificationRepo.DeleteVeruficationByUserId(fmt.Sprintf("%v", userId))
	if err != nil {
		return err
	}

	objectId, _ := primitive.ObjectIDFromHex(fmt.Sprintf("%v", userId))
	userDoc, err := s.FindUserById(objectId)
	if err != nil || userDoc == nil {
		return err
	}

	userDoc.IsEmailVerified = true

	err = s.userRepo.UpdateUserById(userDoc)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(user *models.UserLogin, c *gin.Context) (*models.User, error) {
	userDoc, err := s.FindUserByEmail(user.Email)

	if err != nil {
		return nil, err
	} else if userDoc == nil {
		return nil, errors.New("email or password is not valid")
	}

	err = utils.ComparePassword(userDoc.Password, user.Password)

	if err != nil {
		fmt.Println("userlogin: ", err.Error())

		return nil, errors.New("email or password is not valid")
	}

	if !userDoc.IsApproved {
		return nil, errors.New("account still not approved. please contact with the authority")
	}

	expires := time.Now().Add(time.Minute * 30).Unix()
	token, _ := utils.GenerateJWT(userDoc, "user", expires)
	c.SetCookie("user-token", token, 3600, "/", "", false, true)
	return userDoc, nil
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
	jwt, _ := utils.GenerateJWT(userDoc.ID, "userId", expires)
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

func (s *UserService) FindUserById(id primitive.ObjectID) (*models.User, error) {
	return s.userRepo.GetUserById(id)
}
