package main

import (
	"IbtService/internal/middlware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	r := gin.New()

	r.Use(middlware.RequestLogger(app.log))

	v1 := r.Group("/api")
	{
		v1.POST("/test", app.TestHandler)
		v1.GET("/insertOutBox", app.selectOutBoxCredit)
	}

	return r
}
