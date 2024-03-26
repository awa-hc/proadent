package web

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"services/internal/application/port"

	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func SendContactForm(emailSender port.EmailSender) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req EmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		htmlSender := `
		<!DOCTYPE html>
		<html lang="es">

		<head>
		<meta charset="UTF-8">
		<title>Respuesta Contacto</title>
		<style>
			body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			margin: 0;
			padding: 0;
			}

			.container {
			max-width: 600px;
			margin: 20px auto;
			padding: 20px;
			background-color: #fff;
			border-radius: 5px;
			box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
			}

			h1 {
			color: #333;
			text-align: center;
			}

			p {
			color: #555;
			line-height: 1.6;
			}

			ul {
			padding-left: 20px;
			}
		</style>
		</head>

		<body>

		<div class="container">
			<h1>Gracias por ponerte en contacto con Proadent</h1>
			<p>Hemos recibido tu mensaje y nos pondremos en contacto contigo lo antes posible.</p>
			<p>Detalles del mensaje:</p>
			<ul>
			<li><strong>Email:</strong> ` + req.Email + `</li>
			<li><strong>Teléfono:</strong> ` + req.Phone + `</li>
			<li><strong>Mensaje:</strong> ` + req.Message + `</li>
			</ul>
			<p>Atentamente,<br/>Equipo de Proadent</p>
		</div>

		</body>

		</html>

	`

		htmlAdmin := `
		<!DOCTYPE html>
		<html lang="en">
		
		<head>
		  <meta charset="UTF-8">
		  <meta name="viewport" content="width=device-width, initial-scale=1.0">
		  <title>Contact Form</title>
		  <style>
			body {
			  font-family: Arial, sans-serif;
			  background-color: #f4f4f4;
			  margin: 0;
			  padding: 0;
			}
		
			.container {
			  max-width: 600px;
			  margin: 20px auto;
			  padding: 20px;
			  background-color: #fff;
			  border-radius: 5px;
			  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
			}
		
			h1 {
			  color: #333;
			  text-align: center;
			}
		
			p {
			  color: #555;
			  line-height: 1.6;
			}
		
			.user-email {
			  font-weight: bold;
			  color: #007bff;
			}
		
			.btn-whatsapp {
			  display: block;
			  width: 100%;
			  max-width: 200px;
			  margin: 20px auto;
			  padding: 10px 20px;
			  background-color: #25D366;
			  color: #fff;
			  text-align: center;
			  text-decoration: none;
			  border-radius: 5px;
			}
		
			.message {
			  margin-top: 20px;
			}
		  </style>
		</head>
		
		<body>
		
		  <div class="container">
			<h1>PROADENT CONTACT FORM</h1>
			<p>User with email <span class="user-email">` + req.Email + `</span> has submitted a contact form.</p>
		
			<a href="https://wa.me/` + req.Phone + `" class="btn-whatsapp" target="_blank">Send Message via WhatsApp</a>
		
			<div class="message">
			  <h3>Message:</h3>
			  <p>` + req.Message + `</p>
			</div>
		
			<p>You can also reply via email: <a href="mailto:` + req.Email + `">` + req.Email + `</a></p>
		  </div>
		
		</body>
		
		</html>
		
		
		`

		tmpluser, err := template.New("email").Parse(htmlSender)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing html"})
			return
		}
		tmpladmin, err := template.New("email").Parse(htmlAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing template"})
			return
		}

		var htmlBody bytes.Buffer
		var htmlBodyAdmin bytes.Buffer

		if err := tmpluser.Execute(&htmlBody, nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		if err := tmpladmin.Execute(&htmlBodyAdmin, nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		if err := emailSender.SendEmail(req.Email, "formulario de contacto enviado", htmlBody.String()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		if err := emailSender.SendEmail(os.Getenv("EMAIL_ADMIN"), "formulario de contacto recibido", htmlBodyAdmin.String()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully"})

	}
}
