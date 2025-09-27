package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ProxyUsername string
	ProxyPassword string
	ProxyHost     string
	LogFile       string
	Port          int
}

func Load() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Файл .env не найден прочитайте README.md")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("PORT должен быть числом")
	}

	return &Config{
		ProxyUsername: os.Getenv("PROXY_USERNAME"),
		ProxyPassword: os.Getenv("PROXY_PASSWORD"),
		ProxyHost:     os.Getenv("PROXY_HOST"),
		LogFile:       os.Getenv("LOG_FILE"),
		Port:          port,
	}
}
