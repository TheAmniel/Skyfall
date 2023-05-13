package routes

import (
	"github.com/gofiber/fiber/v2"

	"skyfall/services/config"
	"skyfall/services/database"
)

func Configure(app *fiber.App, db *database.Database, cfg *config.ServerConfig) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, world!"})
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("https://github.com/TheAmniel")
	})
}
