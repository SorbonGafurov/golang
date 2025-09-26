package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ProxyUsername string
	ProxyPassword string
	ProxyHost     string
	LogFile       string
}

func Load() *Config {
	_ = godotenv.Load("../../.env")
	return &Config{
		ProxyUsername: os.Getenv("PROXY_USERNAME"),
		ProxyPassword: os.Getenv("PROXY_PASSWORD"),
		ProxyHost:     os.Getenv("PROXY_HOST"),
		LogFile:       "app.log",
	}
}
