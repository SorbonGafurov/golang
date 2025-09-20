package httpdelivery

import (
	"IbtService/internal/domain"
	"IbtService/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(s service.ExternalService) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqData := &domain.Request{}
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
