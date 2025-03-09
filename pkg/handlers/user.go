package handlers

import (
	"fmt"
	"net/http"

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

func (h *UserHandler) Login(c *gin.Context) {
	var user models.UserLogin

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	if err := user.UserLoginValidate(); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	adminDoc, err := h.service.Login(&user, c)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	utils.SuccessJSON(c, "successfully logged in", http.StatusOK, adminDoc)
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	if err := user.UserRegistrationValidate(); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterUser(c, &user); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.SuccessJSON(c, "successfully registered. an email will sent to your mail for approval", http.StatusCreated, nil)
	// c.JSON(http.StatusCreated, user)
}
