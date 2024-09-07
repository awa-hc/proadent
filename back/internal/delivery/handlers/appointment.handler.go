package handlers

import (
	"back/internal/domain/entities"
	"back/internal/domain/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	AppointmentHandler services.AppointmentService
}

func NewAppointmentHandler(appointmentService services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		AppointmentHandler: appointmentService,
	}
}
func (h *AppointmentHandler) Create(c *gin.Context) {

	var appointment entities.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	if err := h.AppointmentHandler.Created(c, &appointment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appointment)

}
func (h *AppointmentHandler) GetByID(c *gin.Context) {
	codestr := c.Param("id")
	code, err := strconv.Atoi(codestr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	appointment, err := h.AppointmentHandler.GetByID(c, code)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, appointment)
}

func (h *AppointmentHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")

	appointment, err := h.AppointmentHandler.GetByCode(c, code)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appointment)
}

func (h *AppointmentHandler) GetByDoctorCI(c *gin.Context) {
	doctorCI := c.Param("doctorCI")

	appointments, err := h.AppointmentHandler.GetByDoctorCI(c, doctorCI)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appointments)
}

func (h *AppointmentHandler) GetByPatientCI(c *gin.Context) {
	patientCI := c.Param("patientCI")

	appointments, err := h.AppointmentHandler.GetByPatientCI(c, patientCI)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appointments)
}

func (h *AppointmentHandler) GetAll(c *gin.Context) {
	appointments, err := h.AppointmentHandler.GetAll(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, appointments)
}

func (h *AppointmentHandler) Confirm(c *gin.Context) {
	code := c.Param("code")

	if err := h.AppointmentHandler.Confirm(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Appointment confirmed"})
}

func (h *AppointmentHandler) Cancel(c *gin.Context) {
	code := c.Param("code")

	if err := h.AppointmentHandler.Cancel(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Appointment canceled"})
}

func (h *AppointmentHandler) Finish(c *gin.Context) {
	code := c.Param("code")

	if err := h.AppointmentHandler.Finish(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Appointment finished"})
}

func (h *AppointmentHandler) Accept(c *gin.Context) {
	code := c.Param("code")

	if err := h.AppointmentHandler.Accept(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Appointment accepted"})
}

func (h *AppointmentHandler) Reject(c *gin.Context) {
	code := c.Param("code")

	if err := h.AppointmentHandler.Reject(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Appointment rejected"})
}

func (h *AppointmentHandler) Delete(c *gin.Context) {
	code := c.Param("code")

	if err := h.AppointmentHandler.Delete(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Appointment deleted"})
}

func (h *AppointmentHandler) UpdateDoctorCI(c *gin.Context) {
	code := c.Param("code")
	doctorCI := c.Param("doctorCI")

	if err := h.AppointmentHandler.UpdateDoctorCI(c, code, doctorCI); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "DoctorCI updated"})
}

func (h *AppointmentHandler) UpdatePatientCI(c *gin.Context) {
	code := c.Param("code")
	patientCI := c.Param("patientCI")

	if err := h.AppointmentHandler.UpdatePatientCI(c, code, patientCI); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "PatientCI updated"})
}

func (h *AppointmentHandler) UpdateDateTime(c *gin.Context) {
	code := c.Param("code")
	dateTime := c.Param("dateTime")

	if err := h.AppointmentHandler.UpdateDateTime(c, code, dateTime); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "DateTime updated"})
}

func (h *AppointmentHandler) UpdateReason(c *gin.Context) {
	code := c.Param("code")
	var Json struct {
		Reason string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&Json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.AppointmentHandler.UpdateReason(c, code, Json.Reason); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Reason updated"})
}

func (h *AppointmentHandler) UpdateStatus(c *gin.Context) {
	code := c.Param("code")
	status := c.Param("status")

	if err := h.AppointmentHandler.UpdateStatus(c, code, status); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Status updated"})
}

func (h *AppointmentHandler) UpdatePrice(c *gin.Context) {
	code := c.Param("code")
	var Json struct {
		Price float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&Json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.AppointmentHandler.UpdatePrice(c, code, Json.Price); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Price updated"})
}

func (h *AppointmentHandler) UpdateType(c *gin.Context) {
	code := c.Param("code")
	var Json struct {
		Type string `json:"type"`
	}
	if err := c.ShouldBindJSON(&Json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.AppointmentHandler.UpdateType(c, code, Json.Type); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Type updated"})
}
