package controller

import (
	"go-htmx-template/views"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	prefix = ""
)

func AppRoute(route fiber.Router) {
	prefix = viper.GetString("app_prefix")
	views.BaseURL = prefix
	slog.Info("Set Base Route", "prefix", prefix)

	// Static file route
	route.Static(prefix+"/assets", "./static", GetStaticConfig())

	// Rendar page with templ
	// Content-Type: HTML
	r := route.Group(prefix, setDefaultHeader)

	// Index route
	r.Get("/", IndexPage)

	// Components route
	r.Get("/components/*", GetComponents)

}

func setDefaultHeader(c *fiber.Ctx) error {
	// set default header
	// 1. Content-Type: HTML
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.Next()
}
