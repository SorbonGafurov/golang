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
	rabb    *service.Rabbit
	ob      *service.OutBox
}

func main() {

	//конфигурация
	cfgLoad := config.Load()

	//OutBox
	o, err := service.OutBoxOpen(cfgLoad)
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()

	o.InsertOutBox("test message")
	//RabbitMq
	/*r, err := service.ConnRabbit(cfgLoad)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()*/

	//httClient
	client := httpclient.NewProxyClient(cfgLoad)

	app := &application{
		service: service.NewExternalService(client, cfgLoad),
		cfg:     cfgLoad,
		log:     logger.NewLogger(),
		//rabb:    r,
		ob: o,
	}

	err = app.serve()

	if err != nil {
		log.Fatal(err)
	}

}
