package helpers

import (
	"fmt"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	log "log"
	"os"
)

func SendEmail(name, email, message string) {
	fmt.Println(os.Getenv("MAILJET_PUBLIC_API_KEY"))
	fmt.Println(os.Getenv("MAILJET_PRIVATE_API_KEY"))
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MAILJET_PUBLIC_API_KEY"), os.Getenv("MAILJET_PRIVATE_API_KEY"))
	messagesInfo := []mailjet.InfoMessagesV31 {
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "lionel.ritchie@yahoo.com",
				Name: "Lionel",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31 {
					Email: email,
					Name: name,
				},
			},
			Subject: "Greetings from Mailjet.",
			TextPart: message,
			HTMLPart: "<h3> " + message + " </h3>",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}