package httpclient

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"IbtService/internal/config"
)

func NewProxyClient(cfg *config.Config) *http.Client {
	proxyURL, _ := url.Parse(fmt.Sprintf(
		"http://%s:%s@%s",
		cfg.ProxyUsername,
		cfg.ProxyPassword,
		cfg.ProxyHost,
	))

	return &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		Timeout:   30 * time.Second,
	}
}
