package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID       uint16         `gorm:"primaryKey" json:"id"`
	Name     string         `json:"name"`
	Amount   uint32         `json:"amount"`
	Price    uint64         `json:"price"`
	CreateAt *time.Time     `gorm:"default:NOW()" json:"-"`
	UpdateAt *time.Time     `gorm:"default:NOW()" json:"-"`
	DeleteAt gorm.DeletedAt `json:"-"`

	Mart []Mart
}

type ProductUseCase interface {
	CreateProduct(*Product) error
	GetProduct(id uint16)
}

type ProductRepository interface {
}
