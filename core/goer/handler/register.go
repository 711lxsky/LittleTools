package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goer/config"
	MyErr "goer/error"
	"goer/model"
	"goer/reqeust"
	"net/http"
)

// UserRegister 用户注册函数
// TODO 前端传值改变为特定加密 --> 自定义加密方式或者直接使用类SM3的哈希加密
func UserRegister(c *gin.Context) {
	// 尝试解析注册请求
	var RURequest reqeust.RegisterUserRequest
	if err := c.ShouldBindJSON(&RURequest); err != nil {
		// 解析失败
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				config.ResponseMessage: MyErr.JSONParseError,
				config.ResponseReason:  err.Error(),
			})
		return
	}
	// 解析成功，先检查两个密码是否一致
	if RURequest.ConfirmPassword != RURequest.RawPassword {
		// 不一致
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				config.ResponseMessage: MyErr.ConfirmUserPassWordNotSame,
			})
		return
	}
	// 检查用户名是否符合要求
	// TODO 这里的检查逻辑可以进行优化，包括查询判空、空结构体等
	if err := checkUserName(RURequest.UserName); !errors.Is(err, MyErr.CommonError{}) {
		c.AbortWithStatusJSON(
			int(err.Code()),
			gin.H{
				config.ResponseMessage: err.Msg(),
			})
		return
	}
	// 检查通过保存信息
}

func checkUserName(registerUserName string) MyErr.CommonError {
	// 检查长度
	n := len(registerUserName)
	if n < config.UserNameMinLength || n > config.UserNameMaxLength {
		// 长度不符合要求
		return MyErr.NewErrorWithoutReason(http.StatusBadRequest, MyErr.RegisterUserNameLengthError)
	}
	// 检查是否重复
	var user model.User
	repeatNameUserQR := config.DataBase.Where("user_name = ?", registerUserName).First(&user)
	if errors.Is(repeatNameUserQR.Error, gorm.ErrRecordNotFound) {
		return MyErr.CommonError{}
	} else {
		// 查询失败
		return MyErr.NewError(http.StatusBadRequest, MyErr.DataBaseQueryError, repeatNameUserQR.Error.Error())
	}
	// 拿到查询结果
}
