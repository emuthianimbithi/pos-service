package services

import (
	"fmt"
	"log"
)

// NotificationProvider defines the interface for sending notifications
type NotificationProvider interface {
	SendSMS(to, message string) error
	SendEmail(to, subject, body string) error
}

// ConsoleNotificationProvider logs notifications to console (for development)
type ConsoleNotificationProvider struct{}

func NewConsoleNotificationProvider() *ConsoleNotificationProvider {
	return &ConsoleNotificationProvider{}
}

func (p *ConsoleNotificationProvider) SendSMS(to, message string) error {
	log.Printf("[SMS] To: %s | Message: %s", to, message)
	return nil
}

func (p *ConsoleNotificationProvider) SendEmail(to, subject, body string) error {
	log.Printf("[EMAIL] To: %s | Subject: %s | Body: %s", to, subject, body)
	return nil
}

// NotificationService is the main service that uses a provider
type NotificationService struct {
	provider NotificationProvider
}

func NewNotificationService(providerType string) *NotificationService {
	var provider NotificationProvider

	switch providerType {
	case "console":
		provider = NewConsoleNotificationProvider()
	// Future implementations:
	// case "twilio":
	// 	provider = NewTwilioProvider(...)
	// case "sendgrid":
	// 	provider = NewSendGridProvider(...)
	default:
		provider = NewConsoleNotificationProvider()
	}

	return &NotificationService{provider: provider}
}

func (s *NotificationService) SendSMS(to, message string) error {
	if to == "" || message == "" {
		return fmt.Errorf("recipient and message are required")
	}
	return s.provider.SendSMS(to, message)
}

func (s *NotificationService) SendEmail(to, subject, body string) error {
	if to == "" || subject == "" || body == "" {
		return fmt.Errorf("recipient, subject, and body are required")
	}
	return s.provider.SendEmail(to, subject, body)
}

// SendReceipt sends a digital receipt via SMS or Email
func (s *NotificationService) SendReceipt(method, recipient, receiptURL string, amount float64) error {
	message := fmt.Sprintf("Payment of KES %.2f received. View receipt: %s", amount, receiptURL)

	if method == "sms" {
		return s.SendSMS(recipient, message)
	} else if method == "email" {
		subject := "Your Receipt from POS"
		return s.SendEmail(recipient, subject, message)
	}

	return fmt.Errorf("unsupported notification method: %s", method)
}
