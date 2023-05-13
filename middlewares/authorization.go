package middlewares

import "github.com/gofiber/fiber/v2"

func Authorization(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing authorization",
			})
		} else if auth != secret {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	}
}
