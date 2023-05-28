package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
)

func Cache() fiber.Handler {
	return cache.New(cache.Config{
		CacheHeader:  "X-Cache-Status",
		Expiration:   4 * time.Hour,
		CacheControl: true,
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.OriginalURL())
		},
	})
}
