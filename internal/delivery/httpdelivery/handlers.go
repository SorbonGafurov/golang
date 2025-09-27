package httpdelivery

import (
	"IbtService/internal/delivery/httpdelivery/dto"
	"IbtService/internal/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(s service.ExternalService) gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Debug("test")
		reqData := &dto.Request{}

		if err := c.ShouldBindJSON(reqData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		respData, err := s.Send(reqData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, respData)
	}
}
