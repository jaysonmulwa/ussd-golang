package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {

		c.Set("Content-Type", "text/plain")
		return c.SendString("CON Sway the Jay Way Today, Ay?.")
	})

	app.Listen(":3000")
}
