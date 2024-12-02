package handler

import (
	"github.com/gin-gonic/gin"
	"goer/config"
	"net/http"
)

func ResponseFail(c *gin.Context, code int, message string, reason string) {
	c.AbortWithStatusJSON(
		code,
		gin.H{
			config.ResponseMessage: message,
			config.ResponseReason:  reason,
		})
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			config.ResponseMessage: config.Success,
		})
}

func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			config.ResponseMessage: config.Success,
			config.ResponseData:    data,
		})
}
