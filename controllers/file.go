package controllers

import (
	"bytes"
	"io"

	"github.com/gofiber/fiber/v2"
	"skyfall/services/database"
	"skyfall/services/image"
	"skyfall/services/video"
	"skyfall/types"
	"skyfall/utils"
)

// GET /f
func GetAllFile(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var files []types.File
		if err := db.Order("created_at DESC").Find(&files).Error; err != nil {
			return err
		}
		return c.Status(200).JSON(files)
	}
}

// GET /f/{id}
func GetFile(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, ext := utils.ParseFilename(c.Params("id"))
		if !utils.IsImage(ext) && !utils.IsVideo(ext) {
			return fiber.NewError(400, "Invalid extension")
		}

		var file types.File
		if err := db.Select("type", "data").Where("id = ? AND type = ?", id, ext).First(&file).Error; err != nil {
			if database.IsNotFound(err) {
				return fiber.NewError(404, "File not found")
			}
			return err
		}

		if c.QueryBool("thumbnail") {
			if utils.IsVideo(file.Type) {
				if video, err := video.New(file.Data).Thumbnail().Process(); err == nil {
					return c.Status(200).Type("jpeg").SendStream(bytes.NewReader(video), len(video))
				}
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
			return fiber.NewError(400, "Missing file")
		}

		_, extension := utils.ParseFilename(form.Filename)
		if !utils.IsImage(extension) && !utils.IsVideo(extension) && !utils.SupportedMediaType(form.Header["Content-Type"][0]) {
			return fiber.NewError(415, "Unsupported media type")
		}

		file, err := form.Open()
		if err != nil {
			return fiber.NewError(400, "Unable to open file")
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return fiber.NewError(500, "Unable to read file")
		}

		mfile := types.File{Type: extension, Data: data}
		if err := db.Create(&mfile).Error; err != nil {
			return err
		}
		return c.Status(200).JSON(mfile)
	}
}

// DELETE /f/{id}
func DeleteFile(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, ext := utils.ParseFilename(c.Params("id"))
		if !utils.IsImage(ext) && !utils.IsVideo(ext) {
			return fiber.NewError(400, "Invalid extension")
		}

		var file types.File
		if err := db.Select("id", "type").Where("id = ? AND type = ?", id, ext).First(&file).Error; err != nil {
			if database.IsNotFound(err) {
				return fiber.NewError(404, "File not found")
			}
			return err
		}

		if err := db.Unscoped().Delete(&file).Error; err != nil {
			return err
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "OK",
		})
	}
}
