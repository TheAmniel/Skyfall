package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"skyfall/services/config"
	"skyfall/services/database"
)

func Configure(app *fiber.App, db *database.Database, cfg *config.Config) {
	if cfg.Middleware.Logger {
		app.Use(logger.New(logger.Config{
			TimeZone: cfg.Server.TimeZone,
		}))
	}

	if cfg.Middleware.Compress {
		app.Use(compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}))
	}

	if cfg.Middleware.Recover {
		app.Use(recover.New())
	}

	if cfg.Middleware.Banned {
		app.Use(Banned(db))
	}

	if cfg.Middleware.Traffic {
		app.Use(Traffic(db))
	}

	if cfg.Middleware.Cache {
		app.Use(Cache())
	}

	if cfg.Middleware.Shortener {
		app.Use(Shortener(db))
	}
}
