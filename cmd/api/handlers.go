package main

import (
	"IbtService/internal/model"
	"bytes"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) TestHandler(c *gin.Context) {
	reqData := &model.Request{}

	bodyBytes, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := c.ShouldBindBodyWithXML(reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		app.log.Error(err.Error())
		return
	}

	go func() { //временная
		_, _ = app.rabb.PublishToRabbit(bodyBytes)
	}()

	respData, err := app.service.Send(reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		app.log.Error(err.Error())
		return
	}

	xmlResp := &model.Response{}
	if err := xml.Unmarshal(respData.([]byte), xmlResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		app.log.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, xmlResp)
}
