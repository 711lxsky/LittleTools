package handler

import (
	"github.com/gin-gonic/gin"
	MyErr "goer/error"
	"goer/reqeust"
	"goer/service"
	"goer/util"
	"net/http"
	"strings"
)

func UserUpdateSelfInfo(c *gin.Context) {
	// 获取用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析更新请求
	var UURequest *reqeust.UserUpdateSelfInfoRequest
	if err := c.ShouldBindJSON(&UURequest); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	newUserName := strings.TrimSpace(UURequest.UserName)
	if "" != newUserName {
		// 检查用户名
		if err := checkUserName(UURequest.UserName); err != nil {
			ResponseFail(c, http.StatusBadRequest, err.Error(), "")
			return
		}
	}
	newUserEmail := strings.TrimSpace(UURequest.Email)
	if UURequest.Email != "" {
		// 检查邮箱
		if !util.CheckEmailValid(newUserEmail) {
			ResponseFail(c, http.StatusBadRequest, MyErr.UserEmailError, "")
			return
		}
	}
	// 更新用户信息
	if err := service.UpdateUserInfo(userId, newUserName, newUserEmail); err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseUpdateError, err.Error())
		return
	}
	ResponseSuccess(c)
}
