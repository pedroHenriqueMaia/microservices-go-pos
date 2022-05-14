package main

import (
	"microservices/customer-service-onboarding/internal/container"
	"microservices/customer-service-onboarding/internal/http"

	"gitlab.com/pos-alfa-microservices-go/core/config"
	"gitlab.com/pos-alfa-microservices-go/core/http/server"
	"gitlab.com/pos-alfa-microservices-go/core/log"
)

func main() {
	log.Logger.Info("Starting customer-service-onboarding ...")

	conf, err := config.Start()
	if err != nil {
		log.Logger.Fatal("", err)
	}

	container := container.NewContainer(conf)
	if err := container.Start(); err != nil {
		log.Logger.Fatal(err)
	}

	r := http.NewRouter(container)
	if err := server.NewHttpServer(conf.Server.Port).Start(r); err != nil {
		log.Logger.Fatal("fail on start httpserver", err)
	}
}
