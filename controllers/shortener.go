package controllers

import (
	"github.com/gofiber/fiber/v2"
	"skyfall/services/database"
	"skyfall/types"
	"skyfall/utils"
)

// GET /s
func GetAllShort(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var shorts []types.Short
		if err := db.Order("created_at DESC").Find(&shorts).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(200).JSON(shorts)
	}
}

// GET /s/{id}
func GetShort(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var s types.Short
		if err := db.Where("id = ?", c.Params("id")).First(&s).Error; err != nil {
			if database.IsNotFound(err) {
				return c.Status(404).JSON(fiber.Map{"message": "Short not found"})
			}
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(200).JSON(s)
	}
}

// POST /s/upload
func AddShort(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var sh types.Short
		if err := c.BodyParser(&sh); err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid content from request"})
		}

		if !utils.IsURL(sh.URL) {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid URL"})
		}

		if err := db.Create(&sh).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(200).JSON(sh)
	}
}

// DELETE /s/{id}
func DeleteShort(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var sh types.Short
		if err := db.Where("id = ?", c.Params("id")).First(&sh).Error; err != nil {
			if database.IsNotFound(err) {
				return c.Status(404).JSON(fiber.Map{"message": "Short not found"})
			}
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		if err := db.Unscoped().Delete(&sh).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{"message": "OK"})
	}
}
