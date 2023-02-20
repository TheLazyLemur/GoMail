package gomail

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendEmailWithGmail(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Error(err)
	}

	emailAddr := os.Getenv("emailAddr")
	password := os.Getenv("password")

	sender := NewGamilSender("ValleyOfRobots", emailAddr, password)

	subject := "A Test Email"
	content := `
	<h1>A Test Email</h1>
	<p>This is a test email.</p>
	`

	to := []string{"danrousseau@protonmail.com"}
	attachments := []string{"./go.mod"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachments)
	if err != nil {
		t.Error(err)
	}
}
