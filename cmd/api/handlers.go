package main

import (
	"IbtService/internal/model"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) TestHandler(c *gin.Context) {
	reqData := &model.Request{}

	if err := c.ShouldBindJSON(reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	respData, err := app.service.Send(reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	xmlResp := &model.Response{}
	if err := xml.Unmarshal(respData.([]byte), xmlResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, xmlResp)
}

func (app *application) TestHandler_2(c *gin.Context) {
	reqData := &model.RequestTest{}

	if err := c.ShouldBindJSON(reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	respData, err := app.service.Send(reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	xmlResp := &model.ResponseTest{}
	if err := xml.Unmarshal(respData.([]byte), xmlResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, xmlResp)
}
