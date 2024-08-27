package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kategori struct {
	ID   string `gorm:"type:char(36);primaryKey" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

func (kategori *Kategori) BeforeCreate(tx *gorm.DB) (err error) {
	kategori.ID = uuid.New().String()
	return
}
