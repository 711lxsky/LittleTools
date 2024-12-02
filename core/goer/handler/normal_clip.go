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
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// TODO 这里加入了密码逻辑，还需要进行测试，检查前端能否发出对应格式的请求来，以及是否支持发送与不发送密码之间任意切换

// NormalClipboardUse 解析剪切板请求
func NormalClipboardUse(c *gin.Context) {
	var clip *model.Clip
	// 先检查请求的content-type
	contentType := c.GetHeader(config.HeaderContentType)
	switch {
	case strings.Contains(contentType, config.MultipartForm):
		// 处理 图片 + 密码格式
		file, err := c.FormFile(reqeust.ClipFileName)
		if err != nil {
			ResponseFail(c, http.StatusBadRequest, MyErr.FileParseError, err.Error())
			return
		}
		// 尝试解析密码
		var INCRequest *reqeust.ImageNormalClipRequest
		if errINC := c.ShouldBind(&INCRequest); errINC != nil {
			// 密码参数无法解析
			ResponseFail(c, http.StatusBadRequest, MyErr.DataParseError, errINC.Error())
			return
		}
		// 将上传的文件保存到文件路径中
		newNameFileName := util.GenerateNewNameForFile(file.Filename)
		filaSavePath := filepath.Join(config.ImageDirPath, newNameFileName)
		if err := c.SaveUploadedFile(file, filaSavePath); err != nil {
			ResponseFail(c, http.StatusInternalServerError, MyErr.FileSaveError, err.Error())
			return
		}
		encryptPassword, errEP := util.HashEncrypt(INCRequest.Password)
		if errEP != nil {
			ResponseFail(c, http.StatusInternalServerError, MyErr.PasswordEncryptError, errEP.Error())
			return
		}
		// 将数据保存为一个图片类型记录
		clip = &model.Clip{
			Type:       model.ClipImage,
			Content:    newNameFileName,
			Identifier: util.GenerateRandomString(config.IdentifierLength),
			Password:   encryptPassword,
		}
	case strings.Contains(contentType, config.ApplicationJson):
		// 处理 文本 + 密码格式
		var TNCRequest *reqeust.TextNormalClipRequest
		if err := c.ShouldBindJSON(&TNCRequest); err != nil {
			// 解析失败
			ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
			return
		}
		encryptPassword, errEP := util.HashEncrypt(TNCRequest.Password)
		if errEP != nil {
			ResponseFail(c, http.StatusInternalServerError, MyErr.PasswordEncryptError, errEP.Error())
			return
		}
		// 保存一个文本类型的记录到剪切板数据库中
		clip = &model.Clip{
			Type:       model.ClipText,
			Content:    TNCRequest.Content,
			Identifier: util.GenerateRandomString(config.IdentifierLength),
			Password:   encryptPassword,
		}
	default:
		// 不支持的 Content-Type
		ResponseFail(c, http.StatusUnsupportedMediaType, MyErr.UnsupportedMediaType, "")
		return
	}
	// 将记录保存到数据库
	if err := config.DataBase.Create(clip).Error; err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseSaveError, "")
		return
	}
	// 成功响应
	ResponseSuccessWithData(c, clip.Identifier)
}

// NormalClipboardGet 获取一般剪切板使用数据
func NormalClipboardGet(c *gin.Context) {
	// 获取路由中的identifier参数
	identifier := c.Param("identifier")
	// 拿到密码参数
	var GNCRequest *reqeust.GetNormalClipRequest
	if err := c.ShouldBindJSON(&GNCRequest); err != nil {
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	clip, err := service.GetClipByIdentifier(identifier)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataNotExist, "")
		return
	} else if err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseQueryError, "")
		return
	}
	// 拿到查询结果
	if clip.Password != "" {
		// 如果需要密码，则检查密码是否正确
		if !util.CheckHashValid(clip.Password, GNCRequest.Password) {
			// 密码错误
			ResponseFail(c, http.StatusBadRequest, MyErr.PasswordError, "")
			return
		}
	}
	if clip.Type == model.ClipText {
		// 判断是否为文本
		ResponseSuccessWithData(c, clip.Content)
	} else if clip.Type == model.ClipImage {
		// 否则返回图片url
		c.File(filepath.Join(config.ImageDirPath, clip.Content))
	} else {
		// 类型不支持
		log.Fatal("类型不支持")
	}
}
