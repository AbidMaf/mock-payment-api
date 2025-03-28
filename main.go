package main

import (
	"bank-api/config"
	"bank-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.LoadConfig()
	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Listen(":" + config.ServerPort)
}