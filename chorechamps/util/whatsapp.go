package util

import (
	"ChoreChamps/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendWhatsApp(to string, message string) error {
	client := twilio.NewRestClient(config.WhatsAppAccountSID, config.WhatsAppAuthToken)

	_, err := client.ApiV2010.CreateMessage(&openapi.CreateMessageParams{
		From: config.WhatsAppPhoneNumber,
		To:   to,
		Body: message,
	})

	return err
}