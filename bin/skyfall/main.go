package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"skyfall/middlewares"
	"skyfall/services/config"
	"skyfall/services/database"
	_ "skyfall/services/image"
)

func main() {
	cfg := config.Load()
	db := database.New(cfg.Database)

	loc, locErr := time.LoadLocation(cfg.Server.TimeZone)
	if locErr != nil {
		loc = time.Local
	}
	log.SetFlags(0)
	log.SetPrefix("[" + time.Now().In(loc).Format("15:04:05") + "] ")

	app := fiber.New(fiber.Config{
		AppName:       "Skyfall",
		StrictRouting: cfg.Server.StrictRouting,
		CaseSensitive: cfg.Server.CaseSensitive,
		UnescapePath:  cfg.Server.UnescapePath,
		Prefork:       cfg.Server.Prefork,
		BodyLimit:     cfg.Server.Limit << 20,
	})

	/* --- MIDDLEWARES ---*/
	if cfg.Middleware.Logger {
		app.Use(logger.New(logger.Config{
			TimeZone: loc.String(),
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
		app.Use(middlewares.Banned(db))
	}

	if cfg.Middleware.Cache {
		app.Use(cache.New(cache.Config{
			CacheHeader:  "X-Cache-Status",
			Expiration:   24 * time.Hour,
			CacheControl: true,
		}))
	}

	if !fiber.IsChild() {
		log.Printf("Running Skyfall \"%s:%s\"\n", cfg.Server.Host, cfg.Server.Port)
	}
	go func() {
		if err := app.Listen(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)); err != nil {
			log.Fatal(err)
		}
	}()
	if !fiber.IsChild() {
		log.Println("Press CTRL-C to stop the application")
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	if !fiber.IsChild() {
		log.Println("Shutting down Skyfall connection...")
	}

	if err := app.Shutdown(); err != nil {
		log.Println("There was an error while closing the server")
		log.Printf("%T: %v\n", err, err)
	}
}
