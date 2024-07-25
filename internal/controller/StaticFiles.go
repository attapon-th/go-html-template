package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func GetStaticConfig() fiber.Static {
	return fiber.Static{
		Compress:  !viper.GetBool("DEV_MODE"),
		Browse:    false,
		ByteRange: true,
	}
}
