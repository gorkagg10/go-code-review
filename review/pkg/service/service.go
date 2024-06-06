package service

import (
	"coupon_service/pkg/repository"
	. "coupon_service/pkg/service/entity"
	"fmt"
	"github.com/google/uuid"
)

type IService interface {
	ApplyCoupon(Basket, string) (*Basket, error)
	CreateCoupon(int, string, int) any
	GetCoupons([]string) ([]Coupon, error)
}

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) ApplyCoupon(basket Basket, code string) (b *Basket, e error) {
	coupon, err := service.repository.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if basket.Value > 0 {
		basket.AppliedDiscount = coupon.Discount
		if basket.Value < coupon.MinBasketValue {
			basket.ApplicationSuccessful = false
		} else {
			basket.ApplicationSuccessful = true
		}
		return &basket, nil
	}
	if basket.Value == 0 {
		return
	}

	return nil, fmt.Errorf("Tried to apply discount to negative value")
}

func (service *Service) CreateCoupon(discount int, code string, minBasketValue int) any {
	coupon := Coupon{
		ID:             uuid.NewString(),
		Code:           code,
		Discount:       discount,
		MinBasketValue: minBasketValue,
	}

	if err := service.repository.Save(coupon); err != nil {
		return err
	}
	return nil
}

func (service *Service) GetCoupons(codes []string) ([]Coupon, error) {
	coupons := make([]Coupon, 0, len(codes))
	var e error = nil

	for index, code := range codes {
		coupon, err := service.repository.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, index)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, index)
			}
		}
		coupons = append(coupons, *coupon)
	}

	return coupons, e
}
