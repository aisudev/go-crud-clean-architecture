package repository

import (
	"crud/domain"
	"log"

	"gorm.io/gorm"
)

type martRepository struct {
	db *gorm.DB
}

func NewMartRepository(db *gorm.DB) domain.MartRepository {
	return &martRepository{
		db: db,
	}
}

func (r *martRepository) CreateMart(mart *domain.Mart) error {

	if err := r.db.Create(&mart).Error; err != nil {
		log.Fatal("Create user error: ", err.Error())
		return err
	}

	return nil
}

func (r *martRepository) GetMart(id uint16) (*domain.Mart, error) {

	mart := new(domain.Mart)

	if err := r.db.Where("id = ?", id).Find(&mart).Error; err != nil {
		return nil, err
	}

	return mart, nil
}

func (r *martRepository) GetAllMart() ([]domain.Mart, error) {

	var mart []domain.Mart

	if err := r.db.Find(&mart).Error; err != nil {
		return nil, err
	}

	return mart, nil
}

func (r *martRepository) UpdateMart(id uint16, mart domain.Mart) error {

	if err := r.db.Where("id = ?", id).Updates(mart).Error; err != nil {
		return err
	}

	return nil
}

func (r *martRepository) DeleteMart(id uint16) error {

	var mart domain.Mart

	if err := r.db.Where("id = ?", id).Delete(&mart).Error; err != nil {
		return err
	}

	return nil
}
