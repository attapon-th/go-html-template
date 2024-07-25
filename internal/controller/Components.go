package controller

import (
	"go-htmx-template/views"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func GetComponents(c *fiber.Ctx) error {
	components := strings.Split(c.Params("*", ""), "/")
	if len(components) == 0 {
		return c.SendString("")
	}
	com_name := components[0]

	var com = (templ.Component)(nil)

	switch com_name {
	case "auth-buttons":
		com = views.AuthButtons()
	}

	// Render the component
	if com != nil {
		com.Render(c.Context(), c)
	}

	return nil
}
