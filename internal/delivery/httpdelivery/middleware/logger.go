package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger возвращает middleware для логирования HTTP-запросов
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Даем запросу пройти дальше
		c.Next()

		// После того как обработчик отработал
		latency := time.Since(start)
		slog.Info("HTTP request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency,
			"ip", c.ClientIP(),
		)
	}
}
