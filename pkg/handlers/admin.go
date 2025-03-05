package handlers

import (
	"fmt"
	"net/http"

	models "aas.dev/pkg/models/admin"
	"aas.dev/pkg/services"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type AdminHandler struct {
	service *services.AdminService
}

func NewAdminHandler(service *services.AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) VerifyAccount(c *gin.Context) {
	if err := h.service.VerifyAdmin(c); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
}

func (h *AdminHandler) GetAdmins() (interface{}, error) {
	admins, err := h.service.GetAdminUsers()
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (h *AdminHandler) GetAdminByEmail(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)

	admins, err := h.service.GetAdminUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (h *AdminHandler) GetAdminList(p graphql.ResolveParams) (interface{}, error) {
	admins, err := h.service.GetAdminUsers()
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (h *AdminHandler) Login(c *gin.Context) {
	var admin models.AdminLogin

	if err := c.ShouldBindBodyWithJSON(&admin); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	if err := admin.AdminLoginValidate(); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	adminDoc, err := h.service.LoginAdmin(&admin, c)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
	}

	utils.SuccessJSON(c, "successfully logged in", http.StatusOK, adminDoc)
}

func (h *AdminHandler) RegisterUser(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	if err := admin.AdminValidate(); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterAdmin(c, &admin); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.SuccessJSON(c, "successfully registered. an email will sent to your mail after approval", http.StatusCreated, nil)

}
