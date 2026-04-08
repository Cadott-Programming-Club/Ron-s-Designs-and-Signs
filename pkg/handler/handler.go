package handler

import (
	"github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs/pkg/config"
	"github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs/pkg/database"
	"github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs/pkg/email"

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
	e.GET("/llms.txt", h.LLMsTxt)
	e.GET("/", h.Home)
	e.GET("/services", h.Services)
	e.GET("/gallery", h.Gallery)
	e.GET("/contact", h.Contact)
	e.POST("/contact", h.ContactSubmit)
}
