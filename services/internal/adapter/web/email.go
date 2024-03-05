package web

import (
	"bytes"
	"net/http"
	"services/internal/application/port"

	"html/template"

	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Email string `json:"email" binding:"required"`
	Token string `json:"token" binding:"required"`
}

func NewEmailHandler(emailSender port.EmailSender) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req EmailRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		htmlTemplate := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Document</title>
		</head>

		<style>
			body {
				font-family: Arial, sans-serif;
			}
			h1 {
				color: #333;
			}
			a {
				background-color: #4CAF50;
				border: none;
				color: white;
				padding: 15px 32px;
				text-align: center;
				text-decoration: none;
				display: inline-block;
				font-size: 16px;
				margin: 4px 2px;
				cursor: pointer;
			}
		</style>
		<body>
			<h1>¡Hola!</h1>
			<p>Gracias por registrarte en nuestro sitio. Para completar tu registro, haz clic en el siguiente enlace:</p>
			<p>{{.Token}}</p>
			<a href="http://localhost:8080/verify-email?email={{.Email}}&token={{.Token}}">Verificar Email</a>
		</body>
		
		</html>
	`

		tmpl, err := template.New("email").Parse(htmlTemplate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing email template"})
			return
		}

		var htmlBody bytes.Buffer

		err = tmpl.Execute(&htmlBody, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing email template"})
			return
		}

		if err := emailSender.SendEmail(req.Email, "Asunto del Email", htmlBody.String()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	}
}
