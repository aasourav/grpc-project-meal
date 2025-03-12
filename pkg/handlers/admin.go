package handlers

import (
	"fmt"
	"log"
	"net/http"

	models "aas.dev/pkg/models/admin"
	"aas.dev/pkg/models/types"
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
	req, exists := c.Get("req")
	if !exists {
		utils.ErrorJSON(c, fmt.Errorf("request data not found"), http.StatusBadRequest)
		return
	}

	admin, ok := req.(*models.AdminLogin)
	if !ok {
		utils.ErrorJSON(c, fmt.Errorf("invalid request data"), http.StatusBadRequest)
		return
	}
	adminDoc, err := h.service.LoginAdmin(admin, c)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	utils.SuccessJSON(c, "successfully logged in", http.StatusOK, adminDoc)
}

func (h *AdminHandler) PassowordChange(c *gin.Context) {
	req, exists := c.Get("req")
	if !exists {
		utils.ErrorJSON(c, fmt.Errorf("request data not found"), http.StatusBadRequest)
		return
	}

	admin, ok := req.(*types.ResetPassword)
	if !ok {
		utils.ErrorJSON(c, fmt.Errorf("invalid request data"), http.StatusBadRequest)
		return
	}

	if err := h.service.PassowordChange(c, admin); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.SuccessJSON(c, "password changed successfully", http.StatusCreated, nil)
}

func (h *AdminHandler) RegisterUser(c *gin.Context) {
	req, exists := c.Get("req")
	if !exists {
		utils.ErrorJSON(c, fmt.Errorf("request data not found"), http.StatusBadRequest)
		return
	}

	admin, ok := req.(*models.Admin)
	log.Printf("Type of req: %T\n", req)
	if !ok {
		utils.ErrorJSON(c, fmt.Errorf("invalid request data"), http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterAdmin(c, admin); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.SuccessJSON(c, "successfully registered. an email will sent to your mail for approval", http.StatusCreated, nil)

}
