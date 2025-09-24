package main

import (
	"IbtService/internal/delivery/httpdelivery"
	"IbtService/internal/service"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/natefinch/lumberjack"
)

func main() {
	_ = godotenv.Load("../../.env")
	// Настройки прокси
	proxyUsername := os.Getenv("PROXY_USERNAME")
	proxyPassword := os.Getenv("PROXY_PASSWORD")
	proxyHost := os.Getenv("PROXY_HOST")

	proxyURL, _ := url.Parse(fmt.Sprintf("http://%s:%s@%s", proxyUsername, proxyPassword, proxyHost))
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	// Инициализируем сервис
	service := service.NewExternalService(client)

	// Настраиваем lumberjack (ротация логов)
	logFile := &lumberjack.Logger{
		Filename:   "app.log", // куда писать логи
		MaxSize:    5,         // мегабайты
		MaxBackups: 3,         // хранить последние 3 файла
		MaxAge:     28,        // дни хранения
		Compress:   true,      // gzip
	}

	// Создаем slog с выводом в lumberjack
	logger := slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// Gin роутер
	r := gin.New()

	// Middleware логирования запросов
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		slog.Info("HTTP request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency,
			"ip", c.ClientIP(),
		)
	})

	r.POST("/test", httpdelivery.TestHandler(service))
	r.Run(":8080")
}
