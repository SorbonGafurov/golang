package main

import (
	"IbtService/internal/config"
	"IbtService/internal/delivery/httpdelivery"
	"IbtService/internal/delivery/httpdelivery/middleware"
	"IbtService/internal/httpclient"
	"IbtService/internal/logger"
	"IbtService/internal/service"
	"log/slog"

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

	r.POST("/test", httpdelivery.TestHandler(svc))
	r.Run(":8080")
}
