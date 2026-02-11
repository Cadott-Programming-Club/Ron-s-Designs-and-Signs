package email

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type Service struct {
	apiKey       string
	contactEmail string
	enabled      bool
	httpClient   *http.Client
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
		httpClient:   &http.Client{Timeout: 10 * time.Second},
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

// brevoEmailRequest represents the Brevo Send Transactional Email API payload.
type brevoEmailRequest struct {
	Sender      brevoContact   `json:"sender"`
	To          []brevoContact `json:"to"`
	ReplyTo     *brevoContact  `json:"replyTo,omitempty"`
	Subject     string         `json:"subject"`
	HTMLContent string         `json:"htmlContent"`
}

type brevoContact struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

func (s *Service) SendContactNotification(ctx context.Context, data ContactFormData) error {
	if !s.enabled {
		return fmt.Errorf("email service not configured")
	}

	phone := data.Phone
	if phone == "" {
		phone = "Not provided"
	}

	// Email to the owner with the contact form details
	ownerHTML := fmt.Sprintf(`<html><body>
<h2>New Contact Form Submission</h2>
<p><strong>Name:</strong> %s</p>
<p><strong>Email:</strong> <a href="mailto:%s">%s</a></p>
<p><strong>Phone:</strong> %s</p>
<hr>
<p><strong>Message:</strong></p>
<p>%s</p>
</body></html>`, data.Name, data.Email, data.Email, phone, data.Message)

	ownerPayload := brevoEmailRequest{
		Sender: brevoContact{
			Name:  "Ron's Designs and Signs",
			Email: s.contactEmail,
		},
		To: []brevoContact{
			{Email: s.contactEmail},
		},
		ReplyTo: &brevoContact{
			Name:  data.Name,
			Email: data.Email,
		},
		Subject:     fmt.Sprintf("Contact Form: %s <%s>", data.Name, data.Email),
		HTMLContent: ownerHTML,
	}

	if err := s.sendEmail(ctx, ownerPayload); err != nil {
		return fmt.Errorf("sending owner notification: %w", err)
	}

	slog.Info("contact notification email sent to owner",
		"to", s.contactEmail,
		"from_name", data.Name,
		"from_email", data.Email,
	)

	// Confirmation email to the user
	userHTML := fmt.Sprintf(`<html><body>
<h2>Thank you for contacting Ron's Designs and Signs!</h2>
<p>Hi %s,</p>
<p>We've received your message and will get back to you as soon as possible.</p>
<hr>
<p><strong>Your message:</strong></p>
<p>%s</p>
<hr>
<p>Best regards,<br>Ron's Designs and Signs</p>
</body></html>`, data.Name, data.Message)

	userPayload := brevoEmailRequest{
		Sender: brevoContact{
			Name:  "Ron's Designs and Signs",
			Email: s.contactEmail,
		},
		To: []brevoContact{
			{Name: data.Name, Email: data.Email},
		},
		ReplyTo: &brevoContact{
			Name:  "Ron's Designs and Signs",
			Email: s.contactEmail,
		},
		Subject:     "We received your message - Ron's Designs and Signs",
		HTMLContent: userHTML,
	}

	if err := s.sendEmail(ctx, userPayload); err != nil {
		slog.Error("failed to send confirmation email to user", "error", err, "email", data.Email)
		// Don't return error — the owner already got notified
	} else {
		slog.Info("confirmation email sent to user", "to", data.Email)
	}

	return nil
}

func (s *Service) sendEmail(ctx context.Context, payload brevoEmailRequest) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshalling email payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.brevo.com/v3/smtp/email", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("creating email request: %w", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("api-key", s.apiKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending email via Brevo: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		slog.Error("brevo API error", "status", resp.StatusCode, "body", string(respBody))
		return fmt.Errorf("brevo API returned status %d", resp.StatusCode)
	}

	return nil
}
