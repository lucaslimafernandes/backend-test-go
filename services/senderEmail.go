package services

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail() {

	auth := smtp.PlainAuth("", os.Getenv("SENDER_EMAIL"), os.Getenv("PASSWD_EMAIL"), os.Getenv("SMTPHOST"))

	to := []string{os.Getenv("SENDER_EMAIL")}

	msg := []byte("I dont know")

	err := smtp.SendMail(
		os.Getenv("SMTPPORT"),
		auth,
		os.Getenv("SENDER_EMAIL"),
		to,
		msg,
	)

	if err != nil {
		log.Println("Sending e-mail err:", err)
	}

}
