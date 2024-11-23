package main

import (
	"github.com/gin-gonic/gin"
	"goer/handler"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		// 普通剪切板上传
		v1.POST("/clip-upload", handler.NormalClipboardUse)
		// 获取剪切板内容
		v1.GET("/clip/:identifier", handler.NormalClipboardGet)
	}
}
