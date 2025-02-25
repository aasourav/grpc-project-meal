package services

import (
	"errors"
	"time"

	"aas.dev/pkg/interfaces"
	models "aas.dev/pkg/models/admin"
	"aas.dev/pkg/models/types"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminService struct {
	repo interfaces.AdminRepository
}

func NewAdminService(repo interfaces.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
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

	token, _ := utils.GenerateJWT(adminDoc)
	c.SetCookie("admin-token", token, 3600, "/", "", false, true)
	return adminDoc, nil
}

func (s *AdminService) RegisterAdmin(c *gin.Context, admin *models.Admin) error {
	adminDoc, _ := s.FindAdminByEmail(admin.Email)
	if adminDoc != nil {
		return errors.New("email already exist")
	}
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = admin.CreatedAt
	err := s.repo.CreateAdmin(admin)
	if err != nil {
		return err
	}

	emailVerifyData := types.EmailVerifyTypes{
		Email: admin.Email,
		Name:  admin.Name,
	}

	_, err = NewGeneralService(nil).EmailVerify(c, emailVerifyData)
	if err != nil {
		s.repo.DeleteAdminById(admin.ID)
		return err
	}
	return nil
}

func (s *AdminService) FindAdminByEmail(email string) (*models.Admin, error) {
	return s.repo.GetAdminByEmail(email)
}
