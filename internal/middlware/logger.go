package middlware

import (
	"IbtService/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger(log *logger.AppLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		info := "успешно"

		c.Next()

		if c.Writer.Status() != 200 {
			info = "неуспешно"
		}

		log.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"status":  c.Writer.Status(),
			"latency": time.Since(start).String(),
		}).Info(info)

	}
}
