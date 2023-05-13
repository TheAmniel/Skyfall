package types

import (
	"time"

	"gorm.io/gorm"
)

type Ban struct {
	IP        string    `gorm:"type:varchar(15);primaryKey;unique;not null" json:"ip"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (b *Ban) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now()
	return nil
}
