package helpers

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(toEmail, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_SENDER"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Password Reset Request")
	bodyMessage := fmt.Sprintf("To reset your password, click the following link: <a href='%s/reset-password?token=%s'>Reset Password</a>", os.Getenv("BASE_URL_FE"), token)
	m.SetBody("text/html", bodyMessage)

	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailPort := os.Getenv("MAIL_PORT")
	mailUsername := os.Getenv("MAIL_USERNAME")
	mailHost := os.Getenv("MAIL_HOST")
	portMail, _ := strconv.Atoi(mailPort)

	d := gomail.NewDialer(mailHost, portMail, mailUsername, mailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
