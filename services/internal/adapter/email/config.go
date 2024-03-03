package email

import (
	"fmt"
	"os"
	"strconv"
)

// ConfigureSMTP crea y retorna una nueva instancia de SMTPAdapter con la configuración cargada desde las variables de entorno.
func ConfigureSMTP() *SMTPAdapter {
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("EMAIL_PORT")) // Asegúrate de que EMAIL_PORT sea un entero válido.
	if err != nil {
		fmt.Println("EMAIL_PORT debe ser un número válido")
	}

	smtpUsername := os.Getenv("EMAIL_FROM")
	smtpPassword := os.Getenv("EMAIL_PASSWORD")

	return NewSMTPAdapter(smtpHost, smtpPort, smtpUsername, smtpPassword)
}
