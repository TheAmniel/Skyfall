package types

import (
	"time"

	"gorm.io/gorm"
)

type Ban struct {
	IP   string    `gorm:"type:varchar(15);primaryKey;unique;not null" json:"ip"`
	Date time.Time `gorm:"not null" json:"date"`
}

func (b *Ban) BeforeCreate(tx *gorm.DB) error {
	b.Date = time.Now()
	return nil
}
