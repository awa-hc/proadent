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
	Email    string `json:"email" binding:"required"`
	DateTime string `json:"date" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

func SendEmailAppointmentCreated(emailSender port.EmailSender) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req EmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		parsedDate, err := time.Parse(time.RFC3339, req.DateTime)
		if err != nil {
			fmt.Println("Error parsing date:", err)
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
			<h1>Appointment Created</h1>	
			<p>Dear ` + req.Fullname + `</p>
			<p>Your appointment has been created for ` + formattedDate + `</p>
			<p>Appointment Code: ` + req.Code + `</p>
			<p>Thank you for choosing our service</p>
			<p><a href="https://www.google.com/calendar/render?action=TEMPLATE&text=Proadent+Appointment+` + req.Fullname + `&date=` + req.DateTime + `&details=Code+Appointment%3A+` + req.Code + `+&location=Online+Meeting&sf=true&output=xml" target="_blank" style="background-color:#4CAF50; color:white; padding:10px 15px; text-align:center; text-decoration:none; display:inline-block;">Add to Google Calendar</a></p>

		</body>
		</html>
		`

		tmpl, err := template.New("email").Parse(htmlTemplate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		var htmlBody bytes.Buffer

		if err := tmpl.Execute(&htmlBody, nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		if err := emailSender.SendEmail(req.Email, "Appointment Created", htmlBody.String()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Appointment created email sent"})

	}
}
