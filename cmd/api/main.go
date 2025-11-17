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

	ch := make(chan string, 2)

	// Две горутины, каждая вызывает метод
	for i := 0; i < 2; i++ {
		go func() {
			om, err := o.SelectOutBox()
			if err == nil {
				ch <- om.Message
			} else {
				ch <- "пусто"
			}
		}()
	}

	println("1-" + <-ch)
	println("2-" + <-ch)

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
	}

	err = app.serve()

	if err != nil {
		log.Fatal(err)
	}

}
