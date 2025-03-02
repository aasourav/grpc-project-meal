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

	email, err := utils.VerifyJWT(c.Query("u"), "email")
	if err != nil {
		return err
	}

	verificationData, err := s.verificationRepo.GetVerificationDocByEmail(fmt.Sprintf("%v", email))
	if err != nil {
		return err
	}

	if time.Since(verificationData.CreatedAt) > types.VERIFICATION_EXPIRY_SECONDS*time.Second {
		return errors.New("verification link expired")
	}

	err = s.verificationRepo.DeleteVeruficationByEmail(fmt.Sprintf("%v", email))
	if err != nil {
		return err
	}

	adminDoc, err := s.FindAdminByEmail(fmt.Sprintf("%v", email))
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

	expires := time.Now().Add(time.Second * 120).Unix()
	jwt, _ := utils.GenerateJWT(admin.Email, "email", expires)

	verifyLink := utils.GetBaseURL(c) + fmt.Sprintf("/admins/verify?u=%s", jwt)

	emailVerifyData := types.EmailVerifyTypes{
		Email:           admin.Email,
		VerificaionLink: verifyLink,
		Name:            admin.Name,
	}
	_, err = NewGeneralService(nil).EmailVerify(c, emailVerifyData)
	if err != nil {
		adminDoc, _ = s.FindAdminByEmail(admin.Email)
		s.adminRepo.DeleteAdminById(adminDoc.ID)
		return err
	}

	mailData := &verificationModels.Verification{
		Email:     admin.Email,
		CreatedAt: admin.CreatedAt,
	}
	s.verificationRepo.CreateVerificationRepo(mailData)
	return nil
}

func (s *AdminService) FindAdminByEmail(email string) (*models.Admin, error) {
	return s.adminRepo.GetAdminByEmail(email)
}
