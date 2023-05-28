package routes

import (
	"github.com/gofiber/fiber/v2"

	"skyfall/controllers"
	"skyfall/middlewares"
	"skyfall/services/config"
	"skyfall/services/database"
)

func Configure(app *fiber.App, db *database.Database, cfg *config.Config) {
	middlewares.Configure(app, db, cfg)

	file := app.Group("/f")
	file.Get("/", controllers.GetAllFile(db))
	file.Get("/:id", controllers.GetFile(db))
	file.Post("/upload", middlewares.Authorization(cfg.Server.Secret), controllers.AddFile(db))
	file.Delete("/:id", middlewares.Authorization(cfg.Server.Secret), controllers.DeleteFile(db))

	short := app.Group("/s")
	short.Get("/", controllers.GetAllShort(db))
	short.Get("/:id", controllers.GetShort(db))
	short.Post("/upload", middlewares.Authorization(cfg.Server.Secret), controllers.AddShort(db))
	short.Delete("/:id", middlewares.Authorization(cfg.Server.Secret), controllers.DeleteShort(db))

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Endpoint is not found"})
	})
}
