package main

import (
	"github.com/gin-gonic/gin"
	"goer/handler"
	"goer/middleware"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("")
	{
		// 普通剪切板上传
		v1.POST("/clip-use", handler.NormalClipboardUse)
		// 获取剪切板内容
		v1.GET("/clip/:identifier", handler.NormalClipboardGet)
		// 注册
		v1.POST("/register", handler.UserRegister)
		// 登录
		v1.POST("/login", handler.UserLogin)
		// 用户剪切板记录添加
		v1.POST("user-clip-add", middleware.JwtAuthMiddleware(), handler.AddUserClipboard)
		// 用户剪切板记录分页列举
		v1.POST("user-clip-list", middleware.JwtAuthMiddleware(), handler.PageListUserClips)
		// 用户剪切板记录删除
		v1.POST("user-clip-del", middleware.JwtAuthMiddleware(), handler.DeleteUserClip)
		// 用户更新剪切板内容
		v1.POST("user-clip-update", middleware.JwtAuthMiddleware(), handler.UpdateUserClipContent)
	}
}
