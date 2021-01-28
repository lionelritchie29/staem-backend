package helpers

import (
	"gopkg.in/gomail.v2"
)

func SendOTPMail(targetEmail,accountName, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "store.staempowered@gmail.com")
	m.SetHeader("To", targetEmail)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "<div style=\"width: 60%; margin: 0 auto\">\n  <div style=\"background-color: #17212d; padding: 0.75rem; width: 100%\">\n    <img\n      style=\"width: 10%\"\n      src=\"https://store.cloudflare.steamstatic.com/public/shared/images/header/logo_steam.svg?t=962016\"\n      alt=\"\"\n    />\n  </div>\n\n  <div style=\"background-color: #1d2e41; padding: 1rem 2rem\">\n    <h1 style=\"font-size: 1.5rem; color: #f4f4f4; margin-bottom: 1rem\">\n      Dear <span style=\"font-size: 1.5rem; color: #6fa1bb\">"+ accountName +"</span>\n    </h1>\n    <p style=\"color: #f4f4f4\">Here is the OTP code you need to login</p>\n    <p style=\"color: #6fa1bb; font-size: 1.5rem; font-weight: bold\">" + code + "</p>\n\n    <div style=\"height: 7.5rem; margin: 1rem; background: #17212d\"></div>\n  </div>\n</div>\n")

	d := gomail.NewDialer("smtp.gmail.com", 587, "store.staempowered@gmail.com", "zwhrruqnoyhebgea")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}