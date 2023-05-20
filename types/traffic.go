package types

import "time"

type Traffic struct {
	Month string `gorm:"type:varchar(6);not null;primaryKey;unique" json:"month"`
	Total int64  `gorm:"type:bigint;not null" json:"total"`
}

type Visitor struct {
	Path string    `gorm:"type:varchar(20);not null" json:"path"`
	IP   string    `gorm:"type:varchar(15);not null" json:"ip"`
	Date time.Time `gorm:"not null" json:"date"`
}
