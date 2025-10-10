package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	r := gin.New()

	v1 := r.Group("/api")
	{
		v1.POST("/test", app.TestHandler)
		v1.POST("/test_2", app.TestHandler_2)
	}

	return r
}
