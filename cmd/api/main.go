package main

import (
	"IbtService/internal/config"
	"IbtService/internal/httpclient"
	"IbtService/internal/logger"
	"IbtService/internal/service"
	"log"
)

type application struct {
	service service.ExternalService
	cfg     *config.Config
	log     *logger.AppLogger
}

func main() {

	//конфигурация
	cfgLoad := config.Load()

	//httClient
	client := httpclient.NewProxyClient(cfgLoad)

	app := &application{
		service: service.NewExternalService(client, cfgLoad),
		cfg:     cfgLoad,
		log:     logger.NewLogger(),
	}

	err := app.serve()

	if err != nil {
		log.Fatal(err)
	}
}
