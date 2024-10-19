package service

import (
	"notification-service/types"

	"gopkg.in/gomail.v2"
)

type NotificationService struct {
	smtpConfiguration types.SmtpConfigurations
}

func NewNotificationService(smtpConfiguration types.SmtpConfigurations) *NotificationService {
	return &NotificationService{smtpConfiguration: smtpConfiguration}
}

func (s *NotificationService) Notify(consumerMessage map[string]any) (string, error) {
	email := consumerMessage["email"].(string)
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.smtpConfiguration.Login)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Test")
	msg.SetBody("text/plain", "Emotion Service Has Worked.")

	n := gomail.NewDialer(s.smtpConfiguration.Host, s.smtpConfiguration.Port, s.smtpConfiguration.Login, s.smtpConfiguration.AccessToken)
	if err := n.DialAndSend(msg); err != nil {
		return email, err
	}
	return email, nil
}
