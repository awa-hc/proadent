package handlers

import (
	"back/internal/domain/entities"
	"back/internal/domain/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAppointmentsHandler struct {
	UserAppointmentsService services.UserAppointmentsService
}

func NewUserAppointmentsHandler(userAppointmentsService services.UserAppointmentsService) *UserAppointmentsHandler {
	return &UserAppointmentsHandler{
		UserAppointmentsService: userAppointmentsService,
	}
}

func (h *UserAppointmentsHandler) Created(c *gin.Context) {
	var UserAppointments entities.UserAppointments
	if err := c.ShouldBindJSON(&UserAppointments); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserAppointmentsService.Created(c, &UserAppointments); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, UserAppointments)
}

func (h *UserAppointmentsHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	UserAppointments, err := h.UserAppointmentsService.GetByID(c, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, UserAppointments)
}

func (h *UserAppointmentsHandler) GetByUserCI(c *gin.Context) {
	ci := c.Param("ci")

	UserAppointments, err := h.UserAppointmentsService.GetByUserCI(c, ci)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, UserAppointments)
}

func (h *UserAppointmentsHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserAppointmentsService.Delete(c, id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "UserAppointments deleted"})
}
