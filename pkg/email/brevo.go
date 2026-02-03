package email

import (
	"context"
	"fmt"
	"log/slog"
)

type Service struct {
	apiKey       string
	contactEmail string
	enabled      bool
}

func NewService(apiKey, contactEmail string) *Service {
	if apiKey == "" {
		slog.Info("email service disabled (no BREVO_API_KEY)")
		return &Service{enabled: false, contactEmail: contactEmail}
	}

	slog.Info("email service enabled", "contactEmail", contactEmail)
	return &Service{
		apiKey:       apiKey,
		contactEmail: contactEmail,
		enabled:      true,
	}
}

func (s *Service) IsEnabled() bool {
	return s.enabled
}

type ContactFormData struct {
	Name    string
	Email   string
	Phone   string
	Message string
}

func (s *Service) SendContactNotification(ctx context.Context, data ContactFormData) error {
	if !s.enabled {
		return fmt.Errorf("email service not configured")
	}

	slog.Info("would send contact notification email",
		"to", s.contactEmail,
		"from_name", data.Name,
		"from_email", data.Email,
		"phone", data.Phone,
	)

	return nil
}
