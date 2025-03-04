package services

import (
	"errors"
	"fmt"
	"time"

	"aas.dev/pkg/interfaces"
	models "aas.dev/pkg/models/admin"
	"aas.dev/pkg/models/types"
	verificationModels "aas.dev/pkg/models/verification"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminService struct {
	adminRepo        interfaces.AdminRepository
	verificationRepo interfaces.VerifiactionRepository
}

func NewAdminService(adminRepo interfaces.AdminRepository, verificationRepo interfaces.VerifiactionRepository) *AdminService {
	return &AdminService{adminRepo: adminRepo, verificationRepo: verificationRepo}
}

func (s *AdminService) LoginAdmin(admin *models.AdminLogin, c *gin.Context) (*models.Admin, error) {
	adminDoc, err := s.FindAdminByEmail(admin.Email)
	if err != nil {
		return nil, err
	} else if adminDoc == nil {
		return nil, errors.New("email or password is not valid")
	}

	err = utils.ComparePassword(adminDoc.Password, admin.Password)
	if err != nil {
		return nil, errors.New("email or password is not valid")
	}

	if !adminDoc.IsApproved {
		return nil, errors.New("account still not approved. please contact with the authority")
	}

	expires := time.Now().Add(time.Minute * 30).Unix()
	token, _ := utils.GenerateJWT(adminDoc, "user", expires)
	c.SetCookie("admin-token", token, 3600, "/", "", false, true)
	return adminDoc, nil
}

func (s *AdminService) VerifyAdmin(c *gin.Context) error {
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
	adminDoc, err := s.FindAdminById(objectId)
	if err != nil || adminDoc == nil {
		return err
	}

	adminDoc.IsEmailApproved = true

	err = s.adminRepo.UpdateAdminById(adminDoc)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminService) GetAdminUsers() (*[]models.Admin, error) {
	adminUser, err := s.adminRepo.GetAdmins()
	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

func (s *AdminService) GetAdminUserByEmail(email string) (*models.Admin, error) {
	adminUser, err := s.adminRepo.GetAdminByEmail(email)
	if err != nil {
		return nil, err
	}
	return adminUser, nil
}

func (s *AdminService) RegisterAdmin(c *gin.Context, admin *models.Admin) error {
	adminDoc, _ := s.FindAdminByEmail(admin.Email)
	if adminDoc != nil {
		return errors.New("email already exist")
	}
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = admin.CreatedAt
	admin.Password, _ = utils.HashPassword(admin.Password)
	err := s.adminRepo.CreateAdmin(admin)
	if err != nil {
		return err
	}

	adminDoc, err = s.FindAdminByEmail(admin.Email)
	if err != nil {
		return errors.New("admin user find error: " + err.Error())
	}

	expires := time.Now().Add(time.Second * 120).Unix()
	jwt, _ := utils.GenerateJWT(adminDoc.ID, "userId", expires)

	verifyLink := utils.GetBaseURL(c) + fmt.Sprintf("/admins/verify?u=%s", jwt)

	emailVerifyData := types.EmailVerifyTypes{
		Email:           admin.Email,
		VerificaionLink: verifyLink,
		Name:            admin.Name,
	}
	_, err = NewGeneralService(nil).EmailVerify(c, emailVerifyData)
	if err != nil {
		s.adminRepo.DeleteAdminById(adminDoc.ID)
		return err
	}

	mailData := &verificationModels.Verification{
		Email:     admin.Email,
		CreatedAt: admin.CreatedAt,
		UserId:    adminDoc.ID,
	}

	s.verificationRepo.CreateVerificationRepo(mailData)
	return nil
}

func (s *AdminService) FindAdminByEmail(email string) (*models.Admin, error) {
	return s.adminRepo.GetAdminByEmail(email)
}

func (s *AdminService) FindAdminById(id primitive.ObjectID) (*models.Admin, error) {
	return s.adminRepo.GetAdminById(id)
}
