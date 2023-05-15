package types

import (
	"time"

	"gorm.io/gorm"
	"skyfall/utils"
)

var ShortIDLimit int = 4

type Short struct {
	ID        string    `gorm:"type:varchar(4)" json:"id,omitempty"`
	URL       string    `gorm:"type:varchar(64)" json:"url" xml:"url" form:"url"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
}

func (s *Short) BeforeCreate(tx *gorm.DB) error {
generate:
	id := utils.RandomString(ShortIDLimit)
	r := tx.Select("id").Where("id = ?", id).First(&Short{})
	if r.RowsAffected > 0 {
		goto generate // re-generate ID
	}
	s.ID = id
	return nil
}
