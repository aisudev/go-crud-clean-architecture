package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mart struct {
	ID       uint16         `gorm:"primaryKey" json:"id"`
	Name     string         `gorm:"TYPE:varchar(50)" json:"name"`
	Address  string         `gorm:"TYPE:varchar(256)" json:"address"`
	CreateAt *time.Time     `gorm:"default:NOW()" json:"-"`
	UpdateAt *time.Time     `gorm:"default:NOW()" json:"-"`
	DeleteAt gorm.DeletedAt `json:"-"`
}

type MartUsecase interface {
	CreateMart(*Mart) error
	GetMart(id uint16) (*Mart, error)
	GetAllMart() ([]Mart, error)
	UpdateMart(id uint16, mart Mart) error
	DeleteMart(id uint16) error
}

type MartRepository interface {
	CreateMart(*Mart) error
	GetMart(id uint16) (*Mart, error)
	GetAllMart() ([]Mart, error)
	UpdateMart(id uint16, mart Mart) error
	DeleteMart(id uint16) error
}
