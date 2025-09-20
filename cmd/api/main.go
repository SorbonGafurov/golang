package main

import (
	"IbtService/internal/delivery/httpdelivery"
	"IbtService/internal/service"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// Запускаем сервер
	r := gin.Default()
	r.POST("/test", httpdelivery.TestHandler(service))
	r.Run(":8080")
}
