package handler

import (
	"log/slog"

	"ronsdesigns/internal/email"
	"ronsdesigns/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Contact(c echo.Context) error {
	return pages.Contact("", "").Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) ContactSubmit(c echo.Context) error {
	name := c.FormValue("name")
	emailAddr := c.FormValue("email")
	phone := c.FormValue("phone")
	message := c.FormValue("message")

	if name == "" || emailAddr == "" || message == "" {
		return pages.Contact("error", "Please fill in all required fields.").Render(c.Request().Context(), c.Response().Writer)
	}

	_, err := h.db.Conn.ExecContext(c.Request().Context(),
		"INSERT INTO contact_submissions (name, email, phone, message) VALUES (?, ?, ?, ?)",
		name, emailAddr, phone, message,
	)
	if err != nil {
		slog.Error("failed to save contact submission", "error", err)
		return pages.Contact("error", "Failed to save your message. Please try again.").Render(c.Request().Context(), c.Response().Writer)
	}

	if h.email.IsEnabled() {
		err = h.email.SendContactNotification(c.Request().Context(), email.ContactFormData{
			Name:    name,
			Email:   emailAddr,
			Phone:   phone,
			Message: message,
		})
		if err != nil {
			slog.Error("failed to send contact notification email", "error", err)
		}
	}

	return pages.Contact("success", "Thank you! We'll get back to you soon.").Render(c.Request().Context(), c.Response().Writer)
}
