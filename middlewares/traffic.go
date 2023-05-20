package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"skyfall/services/database"
	"skyfall/types"
)

func trafficPeriod() string {
	year, month, _ := time.Now().Date()
	return fmt.Sprintf("%v/%v", int(month), year)
}

func Traffic(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		file := c.Path()
		go func() {
			// TODO: exclude some path
			trafficP := trafficPeriod()
			traffic := types.Traffic{Month: trafficP}
			if err := db.Where("Month = ?", trafficP).First(&traffic).Error; err != nil {
				if database.IsNotFound(err) {
					db.Create(&traffic)
				} else {
					return
				}
			}
			db.Model(&traffic).Update("total", traffic.Total+1)
			db.Create(&types.Visitor{IP: ip, Path: file, Date: time.Now()})
		}()
		return c.Next()
	}
}
