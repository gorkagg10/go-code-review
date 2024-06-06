package entity

import "coupon_service/pkg/service/entity"

type ApplicationRequest struct {
	Code   string
	Basket entity.Basket
}
