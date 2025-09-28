package main

import (
	"IbtService/internal/config"
	"IbtService/internal/httpclient"
	"IbtService/internal/logger"
	"IbtService/internal/service"
	"log/slog"
)

func main() {
	cfg := config.Load()
	logger := logger.NewLogger(cfg.LogFile)
	slog.SetDefault(logger)

	client := httpclient.NewProxyClient(cfg)
	svc := service.NewExternalService(client)
	serve(svc, cfg)
}
