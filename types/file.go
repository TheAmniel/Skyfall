package types

import (
	"time"

	"gorm.io/gorm"
	"skyfall/utils"
)

var FileIDLimit int = 16

type File struct {
	ID        string    `gorm:"type:varchar(16);primaryKey;unique;not null" json:"id"`
	Type      string    `gorm:"type:varchar(4);not null" json:"type"`
	Data      []byte    `gorm:"type:blob;not null" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (f *File) BeforeCreate(tx *gorm.DB) error {
	bytes, err := utils.Gzip(f.Data)
	if err != nil {
		return err
	}

generate:
	id := utils.RandomString(FileIDLimit)
	r := tx.Select("type").Where("id = ? AND type = ?", id, f.Type).First(&File{})
	if r.RowsAffected > 0 {
		goto generate // re-generate ID
	}
	f.ID = id
	f.Data = bytes
	return nil
}

func (f *File) AfterFind(tx *gorm.DB) error {
	bytes, err := utils.Gunzip(f.Data)
	if err != nil {
		return err
	}
	f.Data = bytes
	return nil
}
