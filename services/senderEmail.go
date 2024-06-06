package services

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(body []byte) {

	auth := smtp.PlainAuth("", os.Getenv("SENDER_EMAIL"), os.Getenv("PASSWD_EMAIL"), os.Getenv("SMTPHOST"))

	to := []string{os.Getenv("SENDER_EMAIL")}

	msg := "From: " + os.Getenv("SENDER_EMAIL") + "\n" +
		"To: " + os.Getenv("SENDER_EMAIL") + "\n" +
		"Subject: " + "SendEmail" + "\n\n" +
		string(body)

	err := smtp.SendMail(
		os.Getenv("SMTPHOST")+":"+os.Getenv("SMTPPORT"),
		auth,
		os.Getenv("SENDER_EMAIL"),
		to,
		[]byte(msg),
	)

	if err != nil {
		log.Println("Sending e-mail err:", err)
	}

}
