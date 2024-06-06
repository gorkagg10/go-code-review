package controller

import (
	"coupon_service/pkg/api/entity"
	"coupon_service/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IController interface {
	Apply(c *gin.Context)
	Create(c *gin.Context)
	Get(c *gin.Context)
}

type Controller struct {
	service service.IService
}

func NewController(service service.IService) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Apply(c *gin.Context) {
	apiReq := entity.ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := controller.service.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (controller *Controller) Create(c *gin.Context) {
	apiReq := entity.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	err := controller.service.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	c.Status(http.StatusOK)
}

func (controller *Controller) Get(c *gin.Context) {
	apiReq := entity.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := controller.service.GetCoupons(apiReq.Codes)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coupons)
}
