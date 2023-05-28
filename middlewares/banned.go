package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"skyfall/services/database"
	"skyfall/types"
)

func Banned(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := db.Select("ip").Where("ip = ?", c.IP()).First(&types.Ban{}).Error; err != nil {
			if database.IsNotFound(err) {
				return c.Next()
			}
			return err
		}
		return fiber.NewError(403, "Forbidden")
	}
}
