package handler

import (
	"github.com/gin-gonic/gin"
	MyErr "goer/error"
	"goer/reqeust"
	"goer/service"
	"goer/util"
	"net/http"
)

func UserLogin(c *gin.Context) {
	var ULReq reqeust.UserLoginRequest
	if err := c.ShouldBindJSON(&ULReq); err != nil {
		// 解析登录信息失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	// 解析成功，检查密码
	userInDB, err := service.GetUserByName(ULReq.UserName)
	if err != nil {
		// 查询失败
		ResponseFail(c, http.StatusBadRequest, MyErr.DataBaseQueryError, err.Error())
		return
	}
	if !util.CheckHashValid(userInDB.Password, ULReq.Password) {
		// 密码错误
		ResponseFail(c, http.StatusBadRequest, MyErr.PasswordError, "")
		return
	}
	// 生成token返回
	token, err := util.GenerateTokenWithUserInfo(userInDB.ID)
	if err != nil {
		ResponseFail(c, http.StatusInternalServerError, err.Error(), "")
		return
	}
	ResponseSuccessWithData(c, token)
}
