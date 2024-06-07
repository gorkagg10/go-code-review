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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	basket, err := controller.service.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, basket)
}

func (controller *Controller) Create(c *gin.Context) {
	apiReq := entity.Coupons{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := controller.service.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusCreated)
}

func (controller *Controller) Get(c *gin.Context) {
	apiReq := entity.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	coupons, err := controller.service.GetCoupons(apiReq.Codes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, coupons)
}
