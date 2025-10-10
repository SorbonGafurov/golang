package main

import (
	"IbtService/internal/config"
	"IbtService/internal/httpclient"
	"IbtService/internal/service"
	"log"
)

type application struct {
	service service.ExternalService
	cfg     *config.Config
}

func main() {

	//конфигурация
	cfgLoad := config.Load()

	//httClient
	client := httpclient.NewProxyClient(cfgLoad)

	app := application{
		service: service.NewExternalService(client, cfgLoad),
		cfg:     cfgLoad,
	}

	err := app.serve()

	if err != nil {
		log.Fatal(err)
	}
}
