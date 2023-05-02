package config

import (
	"os"
	"strconv"
)

var (
	// DBConnection is the database connection string
	DBConnection string

	// WhatsAppAccountSID is the Twilio account SID for WhatsApp integration
	WhatsAppAccountSID string

	// WhatsAppAuthToken is the Twilio auth token for WhatsApp integration
	WhatsAppAuthToken string

	// WhatsAppPhoneNumber is the Twilio phone number for WhatsApp integration
	WhatsAppPhoneNumber string
)

// Init initializes the configuration values
func Init() {
	DBConnection = os.Getenv("DB_CONNECTION")
	WhatsAppAccountSID = os.Getenv("WHATSAPP_ACCOUNT_SID")
	WhatsAppAuthToken = os.Getenv("WHATSAPP_AUTH_TOKEN")
	WhatsAppPhoneNumber = os.Getenv("WHATSAPP_PHONE_NUMBER")
}