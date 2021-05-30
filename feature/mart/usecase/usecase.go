package usecase

import "crud/domain"

type martUsecase struct {
	repo domain.MartRepository
}

func NewMartUsecase(repo domain.MartRepository) domain.MartUsecase {
	return &martUsecase{
		repo: repo,
	}
}

func (r *martUsecase) CreateMart(mart *domain.Mart) error {
	return r.repo.CreateMart(mart)
}

func (r *martUsecase) GetMart(id uint16) (*domain.Mart, error) {
	return r.repo.GetMart(id)
}

func (r *martUsecase) GetAllMart() ([]domain.Mart, error) {
	return r.repo.GetAllMart()
}

func (r *martUsecase) UpdateMart(id uint16, mart domain.Mart) error {
	return r.repo.UpdateMart(id, mart)
}

func (r *martUsecase) DeleteMart(id uint16) error {
	return r.repo.DeleteMart(id)
}
