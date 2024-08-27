package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string   `gorm:"type:char(36);primaryKey" json:"id"`
	Name        string   `gorm:"size:255;not null" json:"name"`
	Description string   `gorm:"type:text" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	KategoriID  string   `gorm:"type:char(36)" json:"kategori_id"`
	Kategori    Kategori `gorm:"foreignKey:KategoriID"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New().String()
	return
}
