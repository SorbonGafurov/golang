package main

import (
	"IbtService/internal/config"
	"IbtService/internal/delivery/httpdelivery"
	"IbtService/internal/delivery/httpdelivery/middleware"
	"IbtService/internal/httpclient"
	"IbtService/internal/logger"
	"IbtService/internal/service"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	logger := logger.NewLogger(cfg.LogFile)
	slog.SetDefault(logger)

	client := httpclient.NewProxyClient(cfg)
	svc := service.NewExternalService(client)

	r := gin.New()
	r.Use(middleware.Logger())

	v1 := r.Group("/api")
	{
		v1.POST("/test", httpdelivery.TestHandler(svc))
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	server.ListenAndServe()
}
