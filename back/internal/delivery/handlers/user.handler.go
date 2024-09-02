package handlers

import (
	"back/internal/domain/entities"
	"back/internal/domain/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var User entities.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	if err := h.UserService.CreateUser(c, &User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, User)

}
func (h *UserHandler) GetByCI(c *gin.Context) {
	ci := c.Param("ci")

	User, err := h.UserService.GetByCI(c, ci)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, User)
}

func (h *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")

	User, err := h.UserService.GetByEmail(c, email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, User)
}

func (h *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")

	User, err := h.UserService.GetByUsername(c, username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, User)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {

	var User entities.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.UpdateUser(c, &User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, User)

}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.DeleteUser(c, id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})

}
