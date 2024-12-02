package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goer/config"
	MyErr "goer/error"
	"goer/model"
	"goer/reqeust"
	"goer/service"
	"goer/util"
	"gorm.io/gorm"
	"html"
	"net/http"
	"strings"
)

// UserRegister 用户注册函数
// TODO 前端传值改变为特定加密 --> 自定义加密方式或者直接使用类SM3的哈希加密
func UserRegister(c *gin.Context) {
	// 尝试解析注册请求
	var RURequest reqeust.UserRegisterRequest
	if err := c.ShouldBindJSON(&RURequest); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	// 解析成功，先检查两个密码是否一致
	if RURequest.ConfirmPassword != RURequest.RawPassword {
		// 不一致
		ResponseFail(c, http.StatusBadRequest, MyErr.ConfirmUserPassWordNotSame, "")
		return
	}
	// 检查用户名是否符合要求
	userName := purgeUserName(RURequest.UserName)
	if err := checkUserName(userName); !errors.Is(err, MyErr.CommonError{}) {
		ResponseFail(c, int(err.Code()), err.Msg(), "")
		return
	}
	// 检查通过保存信息
	hashedPassword, err := util.HashEncrypt(RURequest.RawPassword)
	if err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.PasswordEncryptError, "")
		return
	}
	// 保存到数据库
	newUser := &model.User{
		UserName: userName,
		Password: hashedPassword,
	}
	if err := config.DataBase.Create(newUser).Error; err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseSaveError, "")
		return
	}
	ResponseSuccess(c)
}

func purgeUserName(userName string) string {
	return strings.TrimSpace(html.EscapeString(userName))
}

func checkUserName(registerUserName string) MyErr.CommonError {
	// 检查长度
	n := len(registerUserName)
	if n < config.UserNameMinLength || n > config.UserNameMaxLength {
		// 长度不符合要求
		return MyErr.NewErrorWithoutReason(http.StatusBadRequest, MyErr.RegisterUserNameLengthError)
	}
	// 检查是否重复
	_, err := service.GetUserByName(registerUserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有找到，可以
		return MyErr.CommonError{}
	} else if err != nil {
		// 查询失败
		return MyErr.NewError(http.StatusBadRequest, MyErr.DataBaseQueryError, err.Error())
	} else {
		// 已经有了
		return MyErr.NewErrorWithoutReason(http.StatusBadRequest, MyErr.UserNameExisted)
	}
}
