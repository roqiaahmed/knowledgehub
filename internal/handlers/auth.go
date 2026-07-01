package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roqiaahmed/knowledgehub/internal/dto"
	"github.com/roqiaahmed/knowledgehub/internal/services"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) Register(c *gin.Context) {
	var authDTO dto.RegisterAuthRequest
	if err := c.ShouldBindJSON(&authDTO); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "Invalid request data",
				"details": err.Error()})
		return
	}
	if authDTO.Email == "" || authDTO.FullName == "" || authDTO.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	authResponse, err := h.authService.Register(&authDTO)
	if err != nil {
		if err.Error() == "user already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Internal server error",
				"details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": authResponse,
	})
}
