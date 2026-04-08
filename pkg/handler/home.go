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

func (h *Handler) LLMsTxt(c echo.Context) error {
	content := `# Ron's Designs and Signs

> Custom signs, vinyl graphics, and premium apparel in Cadott, Wisconsin.

Ron's Designs and Signs is a local business specializing in custom signage, vinyl graphics, and personalized apparel. Located in Cadott, Wisconsin, we serve businesses and individuals with professional-quality custom products.

## Services

- Custom Signs: Storefront signs, outdoor banners, yard signs, event signage, channel letters
- Vinyl Graphics: Vehicle wraps, window graphics, wall decals, custom lettering, brand graphics
- Custom Apparel: T-shirts, hoodies, team uniforms, business apparel, event merchandise
- Logo Design: Logo creation, brand identity, logo refinement, vector files, brand guidelines

## Contact

- Phone: (715) 579-8471
- Email: ronsdesigns@hotmail.com
- Contact Form: /contact

## Pages

- Home: /
- Services: /services
- Gallery: /gallery
- Contact: /contact
`
	return c.String(http.StatusOK, content)
}
