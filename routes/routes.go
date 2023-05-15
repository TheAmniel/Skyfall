package routes

import (
	"github.com/gofiber/fiber/v2"

	"skyfall/controllers"
	"skyfall/middlewares"
	"skyfall/services/config"
	"skyfall/services/database"
)

func Configure(app *fiber.App, db *database.Database, cfg *config.ServerConfig) {
	file := app.Group("/f")
	file.Get("/", controllers.GetAllFile(db))
	file.Get("/:id", controllers.GetFile(db))
	file.Post("/upload", middlewares.Authorization(cfg.Secret), controllers.AddFile(db))
	file.Delete("/:id", middlewares.Authorization(cfg.Secret), controllers.DeleteFile(db))

	short := app.Group("/s")
	short.Get("/", controllers.GetAllShort(db))
	short.Get("/:id", controllers.GetShort(db))
	short.Post("/upload", middlewares.Authorization(cfg.Secret), controllers.AddShort(db))
	short.Delete("/:id", middlewares.Authorization(cfg.Secret), controllers.DeleteShort(db))

	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("https://github.com/TheAmniel")
	})
}
