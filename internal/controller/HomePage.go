package controller

import (
	"go-html-template/views"

	"github.com/gofiber/fiber/v2"
)

func IndexPage(c *fiber.Ctx) error {

	b := views.HomePage("Home Page")
	if err := b.Render(c.Context(), c); err != nil {
		return err
	}
	return nil
}
