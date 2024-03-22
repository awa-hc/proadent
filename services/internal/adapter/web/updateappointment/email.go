package web

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"services/internal/application/port"
	"time"

	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Email     string `json:"email" binding:"required"`
	Date      string `json:"date" binding:"required"`
	FullName  string `json:"fullName" binding:"required"`
	Code      string `json:"code" binding:"required"`
	UpdatedAt string `json:"updatedAt" binding:"required"`
	Reason    string `json:"reason" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

func SendEmailAppointmentStatusUpdated(emailSender port.EmailSender) gin.HandlerFunc {

	return func(c *gin.Context) {
		var req EmailRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
			return
		}

		parsedDate, err := time.Parse(time.RFC3339, req.Date)

		if err != nil {
			fmt.Println("erorr parsing date")
			return
		}
		formattedDate := parsedDate.Format("January 2, 2006 at 15:04")

		htmlTemplate := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Appointment Created</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f4f4f4;
					padding: 20px;
				}
				h1 {
					color: #333;
					text-align: center;
				}
				p {
					color: #666;
					font-size: 14px;
					line-height: 1.6;
				}
			</style>

		</head>
		<body>
			<h1>Appointment ` + req.Status + `</h1>	
			<p>Dear ` + req.FullName + `</p>
			<p>Your appointment has been confirmed for ` + formattedDate + `</p>
			<p>Your appointment was confirmed at` + req.UpdatedAt + `</p>
			<p>Appointment Code: ` + req.Code + `</p>
			<p>Thank you for choosing our service</p>
			<p><a href="https://www.google.com/calendar/render?action=TEMPLATE&text=Proadent+Appointment+` + req.FullName + `&date=` + formattedDate + `&details=Code+Appointment%3A+` + req.Code + `+&location=Online+Meeting&sf=true&output=xml" target="_blank" style="background-color:#4CAF50; color:white; padding:10px 15px; text-align:center; text-decoration:none; display:inline-block;">Add to Google Calendar</a></p>

		</body>
		</html>
		`

		tmpl, err := template.New("email").Parse(htmlTemplate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		var htmlbody bytes.Buffer

		if err := tmpl.Execute(&htmlbody, nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		if err := emailSender.SendEmail(req.Email, "Appointment "+req.Status, htmlbody.String()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Update status email sent"})

	}

}
