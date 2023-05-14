package routes

import (
	"github.com/gofiber/fiber/v2"

	"skyfall/controllers"
	"skyfall/middlewares"
	"skyfall/services/config"
	"skyfall/services/database"
)

func Configure(app *fiber.App, db *database.Database, cfg *config.ServerConfig) {
	app.Get("/", controllers.GetAllFile(db))
	app.Get("/:id", controllers.GetFile(db))
	app.Post("/upload", middlewares.Authorization(cfg.Secret), controllers.AddFile(db))
	app.Delete("/:id", middlewares.Authorization(cfg.Secret), controllers.DeleteFile(db))

	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("https://github.com/TheAmniel")
	})
}
