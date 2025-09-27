package main

import (
	"IbtService/internal/config"
	"IbtService/internal/service"
	"fmt"
	"net/http"
	"time"
)

func serve(svc service.ExternalService, cfg *config.Config) error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      route(svc),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server.ListenAndServe()
}
