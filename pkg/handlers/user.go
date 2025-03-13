package handlers

import (
	"fmt"
	"net/http"

	"aas.dev/pkg/models/types"
	models "aas.dev/pkg/models/user"
	"aas.dev/pkg/services"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) VerifyAccount(c *gin.Context) {
	if err := h.service.VerifyUser(c); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
}

func (h *UserHandler) PassowordChange(c *gin.Context) {
	req, exists := c.Get("req")
	if !exists {
		utils.ErrorJSON(c, fmt.Errorf("request data not found"), http.StatusBadRequest)
		return
	}

	user, ok := req.(*types.ResetPassword)
	if !ok {
		utils.ErrorJSON(c, fmt.Errorf("invalid request data"), http.StatusBadRequest)
		return
	}

	if err := h.service.PassowordChange(c, user); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.SuccessJSON(c, "password changed successfully", http.StatusCreated, nil)
}

func (h *UserHandler) Login(c *gin.Context) {
	req, exists := c.Get("req")
	if !exists {
		utils.ErrorJSON(c, fmt.Errorf("request data not found"), http.StatusBadRequest)
		return
	}

	user, ok := req.(*models.UserLogin)
	if !ok {
		utils.ErrorJSON(c, fmt.Errorf("invalid request data"), http.StatusBadRequest)
		return
	}

	adminDoc, err := h.service.Login(user, c)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	utils.SuccessJSON(c, "successfully logged in", http.StatusOK, adminDoc)
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	req, exists := c.Get("req")
	if !exists {
		utils.ErrorJSON(c, fmt.Errorf("request data not found"), http.StatusBadRequest)
		return
	}

	user, ok := req.(*models.User)
	if !ok {
		utils.ErrorJSON(c, fmt.Errorf("invalid request data"), http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterUser(c, user); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.SuccessJSON(c, "successfully registered. an email will sent to your mail for approval", http.StatusCreated, nil)
}
