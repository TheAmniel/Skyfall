package middlewares

import "github.com/gofiber/fiber/v2"

func Authorization(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return fiber.NewError(400, "Missing authorization")
		} else if auth != secret {
			return fiber.NewError(401, "Unauthorized")
		}
		return c.Next()
	}
}
