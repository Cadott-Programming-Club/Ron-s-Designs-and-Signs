package handler

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"ronsdesigns/pkg/config"
	"ronsdesigns/pkg/database"
	"ronsdesigns/pkg/email"
	"ronsdesigns/pkg/handler"
	"ronsdesigns/pkg/middleware"

	"github.com/labstack/echo/v4"
)

var (
	e    *echo.Echo
	once sync.Once
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))
}

func getEcho() *echo.Echo {
	once.Do(func() {
		cfg := config.Load()

		ctx := context.Background()
		db, err := database.New(ctx, cfg.DatabaseURL)
		if err != nil {
			slog.Error("failed to connect to database", "error", err)
			panic(err)
		}

		emailSvc := email.NewService(cfg.BrevoAPIKey, cfg.ContactEmail)

		e = echo.New()
		e.HideBanner = true
		e.HidePort = true

		middleware.Setup(e, cfg)

		h := handler.New(cfg, db, emailSvc)
		h.RegisterRoutes(e)
	})

	return e
}

// Handler is the Vercel serverless function entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	getEcho().ServeHTTP(w, r)
}
