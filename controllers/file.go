package controllers

import (
	"bytes"
	"io"

	"github.com/gofiber/fiber/v2"
	"skyfall/services/database"
	"skyfall/services/image"
	_ "skyfall/services/video"
	"skyfall/types"
	"skyfall/utils"
)

// GET /f
func GetAllFile(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var files []types.File
		if err := db.Order("created_at DESC").Find(&files).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(200).JSON(files)
	}
}

// GET /f/{id}
func GetFile(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, ext := utils.ParseFilename(c.Params("id"))
		if !utils.IsImage(ext) && !utils.IsVideo(ext) {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid extension",
			})
		}

		var file types.File
		if err := db.Select("type", "data").Where("id = ? AND type = ?", id, ext).First(&file).Error; err != nil {
			if database.IsNotFound(err) {
				return c.Status(404).JSON(fiber.Map{
					"message": "File not found",
				})
			}
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if c.QueryBool("thumbnail") {
			if utils.IsVideo(file.Type) {
				// TODO
			} else if utils.IsImage(file.Type) {
				if img, err := image.New(file.Data).Thumbnail().Process(); err == nil {
					return c.Status(200).Type(file.Type).SendStream(bytes.NewReader(img), len(img))
				}
			}
		} else if size := c.QueryInt("size"); size >= 16 && utils.IsImage(file.Type) {
			if img, err := image.New(file.Data).Size(size).Process(); err == nil {
				return c.Status(200).Type(file.Type).SendStream(bytes.NewReader(img), len(img))
			}
		}
		return c.Status(200).Type(file.Type).SendStream(bytes.NewReader(file.Data), len(file.Data))
	}
}

// POST /f/upload
func AddFile(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing file",
			})
		}

		_, extension := utils.ParseFilename(form.Filename)
		if !utils.IsImage(extension) && !utils.IsVideo(extension) && !utils.SupportedMediaType(form.Header["Content-Type"][0]) {
			return c.Status(415).JSON(fiber.Map{
				"message": "Unsupported media type",
			})
		}

		file, err := form.Open()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Unable to open file",
			})
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Unable to read file",
			})
		}

		mfile := types.File{Type: extension, Data: data}
		if err := db.Create(&mfile).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(200).JSON(mfile)
	}
}

// DELETE /f/{id}
func DeleteFile(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, ext := utils.ParseFilename(c.Params("id"))
		if !utils.IsImage(ext) && !utils.IsVideo(ext) {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid extension",
			})
		}

		var file types.File
		if err := db.Select("id", "type").Where("id = ? AND type = ?", id, ext).First(&file).Error; err != nil {
			if database.IsNotFound(err) {
				return c.Status(404).JSON(fiber.Map{
					"message": "File not found",
				})
			}
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := db.Unscoped().Delete(&file).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "OK",
		})
	}
}
