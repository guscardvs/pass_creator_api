package main

import (
	"github.com/gofiber/fiber/v2"
)

func FiberRun() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		passphrase := GetPassPhrase()
		return c.JSON(passphrase)
	})

	app.Listen(":3000")

}
