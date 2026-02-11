package handler

import (
	"net/http"

	"github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) Home(c echo.Context) error {
	return pages.Home().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Services(c echo.Context) error {
	return pages.Services().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Gallery(c echo.Context) error {
	return pages.Gallery().Render(c.Request().Context(), c.Response().Writer)
}
