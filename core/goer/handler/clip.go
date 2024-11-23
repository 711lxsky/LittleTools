package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goer/config"
	MyErr "goer/error"
	"goer/model"
	"goer/reqeust"
	"goer/util"
	"log"
	"net/http"
	"path/filepath"
)

// NormalClipboardUse 解析剪切板请求
func NormalClipboardUse(c *gin.Context) {
	// 先尝试进行文件解析
	file, err := c.FormFile(reqeust.NormalClipFileName)
	var clip *model.Clip
	if errors.Is(err, http.ErrNotMultipart) {
		var normalCR reqeust.NormalClipboardRequest
		if err := c.ShouldBindJSON(&normalCR); err != nil {
			// 解析失败
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					config.ResponseMessage: MyErr.JSONParseError,
					config.ResponseReason:  err.Error(),
				})
			return
		}
		// 否侧保存一个文本类型的记录到剪切板数据库中
		clip = &model.Clip{
			Type:       model.ClipText,
			Content:    normalCR.Content,
			Identifier: util.GenerateRandomString(config.IdentifierLength),
		}
	} else if errors.Is(err, http.ErrMissingFile) {
		// 没有上传文件，则尝试解析文本
		var normalCR reqeust.NormalClipboardRequest
		if err := c.ShouldBindJSON(&normalCR); err != nil {
			// 解析失败
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					config.ResponseMessage: MyErr.JSONParseError,
					config.ResponseReason:  err.Error(),
				})
			return
		}
		// 否侧保存一个文本类型的记录到剪切板数据库中
		clip = &model.Clip{
			Type:       model.ClipText,
			Content:    normalCR.Content,
			Identifier: util.GenerateRandomString(config.IdentifierLength),
		}
	} else if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				config.ResponseMessage: MyErr.FileParseError,
				config.ResponseReason:  err.Error(),
			})
		return
	} else {
		// 将上传的文件保存到文件路径中
		newNameFileName := util.GenerateNewNameForFile(file.Filename)
		filaSavePath := filepath.Join(config.ImageDirPath, newNameFileName)
		if err := c.SaveUploadedFile(file, filaSavePath); err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					config.ResponseMessage: MyErr.FileSaveError,
					config.ResponseReason:  err.Error(),
				})
			return
		}
		// 否则则将数据保存为一个图片类型记录
		clip = &model.Clip{
			Type:       model.ClipImage,
			Content:    newNameFileName,
			Identifier: util.GenerateRandomString(config.IdentifierLength),
		}
	}
	// 将记录保存到数据库
	if err := config.DataBase.Create(clip).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				config.ResponseMessage: MyErr.DataBaseSaveError,
				config.ResponseReason:  err.Error(),
			})
		return
	}
	// 成功相应
	c.JSON(
		http.StatusOK,
		gin.H{
			config.ResponseMessage: config.Success,
			config.ResponseData:    clip.Identifier,
		})
}

// NormalClipboardGet 获取一般剪切板使用数据
func NormalClipboardGet(c *gin.Context) {
	// 获取路由中的identifier参数
	identifier := c.Param("identifier")
	var clip model.Clip
	// 裸写查询语句进行查询
	queryRes := config.DataBase.Where("identifier = ?", identifier).First(&clip)
	if queryRes.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				config.ResponseMessage: MyErr.DataBaseQueryError,
				config.ResponseReason:  queryRes.Error.Error(),
			})
		return
	}
	// 拿到查询结果
	if clip.Type == model.ClipText {
		// 判断是否为文本
		c.JSON(
			http.StatusOK,
			gin.H{
				config.ResponseMessage: config.Success,
				config.ResponseData:    clip.Content,
			})
	} else if clip.Type == model.ClipImage {
		// 否则返回图片url
		c.File(filepath.Join(config.ImageDirPath, clip.Content))
	} else {
		// 类型不支持
		log.Fatal("类型不支持")
	}
}
