package main

import (
	"IbtService/internal/delivery/httpdelivery"
	"IbtService/internal/delivery/httpdelivery/middleware"
	"IbtService/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func route(svc service.ExternalService) http.Handler {
	r := gin.New()
	r.Use(middleware.Logger())

	v1 := r.Group("/api")
	{
		v1.POST("/test", httpdelivery.TestHandler(svc))
	}

	return r
}
