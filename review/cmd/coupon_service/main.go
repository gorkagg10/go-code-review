package main

import (
	"coupon_service/pkg/api"
	"coupon_service/pkg/config"
	control "coupon_service/pkg/controller"
	repo "coupon_service/pkg/repository"
	serv "coupon_service/pkg/service"
	"fmt"
	"time"
)

var (
	configuration = config.NewConfig()
	repository    = repo.NewRepository()
)

func main() {
	service := serv.NewService(repository)
	controller := control.NewController(service)
	server := api.NewApi(configuration.Api, controller)
	server.Start()
	fmt.Println("Starting Coupon service server")
	<-time.After(1 * time.Hour * 24 * 365)
	fmt.Println("Coupon service server alive for a year, closing")
	server.Close()
}
