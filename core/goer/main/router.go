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
		v1.POST("/clip/:identifier", handler.NormalClipboardGet)
		// 注册
		v1.POST("/register", handler.UserRegister)
		// 登录
		v1.POST("/login", handler.UserLogin)
		// 用户更新自身信息
		v1.POST("/user-update", middleware.JwtAuthMiddleware(), handler.UserUpdateSelfInfo)
		// 用户剪切板记录添加
		v1.POST("/user-clip-add", middleware.JwtAuthMiddleware(), handler.AddUserClipboard)
		// 用户剪切板记录分页列举
		v1.POST("/user-clip-list", middleware.JwtAuthMiddleware(), handler.PageListUserClips)
		// 用户剪切板记录删除
		v1.POST("/user-clip-del", middleware.JwtAuthMiddleware(), handler.DeleteUserClip)
		// 用户更新剪切板内容
		v1.POST("/user-clip-update", middleware.JwtAuthMiddleware(), handler.UpdateUserClipContent)
		// 更新剪切板使用时间
		v1.POST("/user-clip-use", middleware.JwtAuthMiddleware(), handler.UpdateUserClipUseTime)
		// 用户待办事项添加
		v1.POST("/todo-add", middleware.JwtAuthMiddleware(), handler.AddTodo)
		// 用户分页查询待办，返回一个简单试图列表
		v1.POST("/todo-list", middleware.JwtAuthMiddleware(), handler.PageTodoList)
		// 用户查询某个待办详情
		v1.POST("/todo-detail", middleware.JwtAuthMiddleware(), handler.GetTodoDetailInfo)
		// 用户删除某个待办
		v1.POST("/todo-del", middleware.JwtAuthMiddleware(), handler.DeleteTodo)
		// 用户更新某个待办
		v1.POST("/todo-update", middleware.JwtAuthMiddleware(), handler.UpdateTodo)
	}
}
