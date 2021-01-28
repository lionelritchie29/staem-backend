package helpers

import (
	"log"
	"net/smtp"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <store.staempowered@gmail.com>"
const CONFIG_AUTH_EMAIL = "store.staempowered@gmail.com"
const CONFIG_AUTH_PASSWORD = "darrowred29"

func SendMail(body string) {
	from := "store.staempowered@gmail.com"
	pass := "zwhrruqnoyhebgea"
	to := "lionelrtchieee@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}