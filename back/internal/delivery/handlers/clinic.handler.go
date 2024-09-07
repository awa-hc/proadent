package handlers

import (
	"back/internal/domain/entities"
	"back/internal/domain/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClinicHandler struct {
	ClinicService services.ClinicService
}

func NewClinicHandler(clinicService services.ClinicService) *ClinicHandler {
	return &ClinicHandler{
		ClinicService: clinicService,
	}
}
func (h *ClinicHandler) Create(c *gin.Context) {

	var clinic entities.Clinic
	if err := c.ShouldBindJSON(&clinic); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	if err := h.ClinicService.Create(c, &clinic); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, clinic)

}

func (h *ClinicHandler) GetByID(c *gin.Context) {
	codestr := c.Param("id")
	codeint, err := strconv.Atoi(codestr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	clinic, err := h.ClinicService.GetByID(c, codeint)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinic)
}

func (h *ClinicHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")

	clinic, err := h.ClinicService.GetByCode(c, code)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinic)
}

func (h *ClinicHandler) GetAll(c *gin.Context) {
	clinics, err := h.ClinicService.GetAll(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinics)
}

func (h *ClinicHandler) GetLast(c *gin.Context) {
	clinic, err := h.ClinicService.GetLast(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinic)
}

func (h *ClinicHandler) Confirm(c *gin.Context) {
	code := c.Param("code")

	if err := h.ClinicService.Confirm(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "confirmed"})
}

func (h *ClinicHandler) Cancel(c *gin.Context) {
	code := c.Param("code")

	if err := h.ClinicService.Cancel(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "canceled"})
}

func (h *ClinicHandler) Finish(c *gin.Context) {
	code := c.Param("code")

	if err := h.ClinicService.Finish(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "finished"})
}

func (h *ClinicHandler) Accept(c *gin.Context) {
	code := c.Param("code")

	if err := h.ClinicService.Accept(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "accepted"})
}

func (h *ClinicHandler) Reject(c *gin.Context) {
	code := c.Param("code")

	if err := h.ClinicService.Reject(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "rejected"})
}

func (h *ClinicHandler) Delete(c *gin.Context) {
	code := c.Param("code")

	if err := h.ClinicService.Delete(c, code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

func (h *ClinicHandler) UpdateDoctorCI(c *gin.Context) {
	code := c.Param("code")
	doctorCI := c.Param("doctorCI")

	if err := h.ClinicService.UpdateDoctorCI(c, code, doctorCI); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "doctorCI updated"})
}

func (h *ClinicHandler) UpdateDateTime(c *gin.Context) {
	code := c.Param("code")
	dateTime := c.Param("dateTime")

	if err := h.ClinicService.UpdateDateTime(c, code, dateTime); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "dateTime updated"})
}

func (h *ClinicHandler) GetLastByUserCI(c *gin.Context) {
	userCI := c.Param("userCI")

	clinic, err := h.ClinicService.GetLastByUserCI(c, userCI)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinic)
}

func (h *ClinicHandler) GetLastByDoctorCI(c *gin.Context) {
	doctorCI := c.Param("doctorCI")

	clinic, err := h.ClinicService.GetLastByDoctorCI(c, doctorCI)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinic)
}

func (h *ClinicHandler) GetByDoctorCI(c *gin.Context) {
	doctorCI := c.Param("doctorCI")

	clinics, err := h.ClinicService.GetByDoctorCI(c, doctorCI)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinics)
}

func (h *ClinicHandler) GetLastByCI(c *gin.Context) {
	ci := c.Param("ci")

	clinic, err := h.ClinicService.GetLastByCI(c, ci)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinic)
}

func (h *ClinicHandler) GetByPatientCI(c *gin.Context) {
	patientCI := c.Param("ci")

	clinics, err := h.ClinicService.GetByPatientCI(c, patientCI)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clinics)
}

func (h *ClinicHandler) UpdateReason(c *gin.Context) {
	code := c.Param("code")
	reason := c.Param("reason")

	if err := h.ClinicService.UpdateReason(c, code, reason); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "reason updated"})
}

func (h *ClinicHandler) UpdateStatus(c *gin.Context) {
	code := c.Param("code")
	status := c.Param("status")

	if err := h.ClinicService.UpdateStatus(c, code, status); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "status updated"})
}

func (h *ClinicHandler) UpdatePrice(c *gin.Context) {
	code := c.Param("code")
	pricestr := c.Param("price")

	price, err := strconv.ParseFloat(pricestr, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.ClinicService.UpdatePrice(c, code, price); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "patientCI updated"})
}
