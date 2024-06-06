package api

import (
	"context"
	"coupon_service/pkg/config"
	"coupon_service/pkg/controller"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct {
	server     *http.Server
	mux        *gin.Engine
	controller controller.IController
	config     config.ApiConfig
}

func NewApi(config *config.ApiConfig, controller controller.IController) Api {
	gin.SetMode(gin.ReleaseMode)
	router := new(gin.Engine)
	router = gin.New()
	router.Use(gin.Recovery())

	return Api{
		mux:        router,
		config:     *config,
		controller: controller,
	}.withServer().withRoutes()
}

func (api Api) withServer() Api {

	channel := make(chan Api)
	go func() {
		api.server = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", api.config.Host, api.config.Port),
			Handler: api.mux,
		}
		channel <- api
	}()

	return <-channel
}

func (api Api) withRoutes() Api {
	apiGroup := api.mux.Group("/api")
	apiGroup.POST("/apply", api.controller.Apply)
	apiGroup.POST("/coupons", api.controller.Create)
	apiGroup.GET("/coupons", api.controller.Get)
	return api
}

func (api Api) Start() {
	if err := api.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (api Api) Close() {
	<-time.After(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := api.server.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
