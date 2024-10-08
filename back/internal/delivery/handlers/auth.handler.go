package handlers

import (
	"back/internal/domain/entities"
	"back/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService services.Auth
}

func NewAuthHandler(authService services.Auth) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) LoginWithCI(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	token, err := h.AuthService.LoginWithCI(c, email, password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")

	err := h.AuthService.Logout(c, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Logout successful"})
}
func (h *AuthHandler) LoginWithEmail(c *gin.Context) {
	var Request entities.Login

	if err := c.ShouldBindJSON(&Request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.LoginWithEmail(c, Request.Email, Request.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", token, 3600*24, "/", "", false, true)

	c.JSON(200, gin.H{"token": token})
}
func (h *AuthHandler) GetUserByContext(c *gin.Context) {
	User, err := h.AuthService.GetUserByContext(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, User)
}
