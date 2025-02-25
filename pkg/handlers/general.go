package handlers

import (
	"fmt"
	"net/http"

	"aas.dev/pkg/models/types"
	"aas.dev/pkg/services"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
)

// GeneralHandler struct for handling general application routes.
type GeneralHandler struct {
	service *services.GeneralService
}

// NewMainHandler creates a new instance of GeneralHandler.
func NewGeneralHandler() *GeneralHandler {
	return &GeneralHandler{}
}

// HealthCheck responds with a simple health status.
func (h *GeneralHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *GeneralHandler) AboutUs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "meal management", "version": "0.0.1", "developer": "ahsan amin", "email": "ahsan.sourav109@gmail.com"})
}

func (h *GeneralHandler) EmailVerify(c *gin.Context) {
	var emailRequestData types.EmailVerifyTypes
	if err := c.ShouldBindBodyWithJSON(&emailRequestData); err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		return
	}

	mailSvc, err := h.service.EmailVerify(c, emailRequestData)
	if err != nil {
		utils.ErrorJSON(c, err, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	utils.SuccessJSON(c, mailSvc.Response, http.StatusOK, nil)
}
