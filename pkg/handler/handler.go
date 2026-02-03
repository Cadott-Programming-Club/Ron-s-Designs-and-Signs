package handler

import (
	"ronsdesigns/pkg/config"
	"ronsdesigns/pkg/database"
	"ronsdesigns/pkg/email"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg   *config.Config
	db    *database.DB
	email *email.Service
}

func New(cfg *config.Config, db *database.DB, emailSvc *email.Service) *Handler {
	return &Handler{
		cfg:   cfg,
		db:    db,
		email: emailSvc,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.Static("/static", "static")
	e.Static("/images", "images")

	e.GET("/health", h.Health)
	e.GET("/", h.Home)
	e.GET("/services", h.Services)
	e.GET("/gallery", h.Gallery)
	e.GET("/contact", h.Contact)
	e.POST("/contact", h.ContactSubmit)
}
