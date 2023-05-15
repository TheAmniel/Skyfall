package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"skyfall/services/database"
	"skyfall/types"
)

func Shortener(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Path()[1:]
		ln := len(id)
		if ln == types.ShortIDLimit {
			var short types.Short
			if err := db.Select("URL").Where("id = ?", id).First(&short).Error; err != nil {
				if database.IsNotFound(err) {
					return c.Next()
				}
				return err
			}
			return c.Redirect(short.URL, 301)
		}
		return c.Next()
	}
}
