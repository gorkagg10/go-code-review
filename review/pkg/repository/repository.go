package repository

import (
	"coupon_service/pkg/service/entity"
	"fmt"
)

type Config struct{}

type IRepository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
}

type Repository struct {
	entries map[string]entity.Coupon
}

func NewRepository() *Repository {
	return &Repository{
		entries: make(map[string]entity.Coupon),
	}
}

func (repository *Repository) FindByCode(code string) (*entity.Coupon, error) {
	coupon, ok := repository.entries[code]
	if !ok {
		return nil, fmt.Errorf("coupon not found")
	}
	return &coupon, nil
}

func (repository *Repository) Save(coupon entity.Coupon) error {
	repository.entries[coupon.Code] = coupon
	return nil
}
