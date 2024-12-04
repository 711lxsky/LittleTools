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
	"net/http"
	"path/filepath"
	"time"
)

func AddUserClipboard(c *gin.Context) {
	// 拿到用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 再解析剪切板请求
	var userClip *model.UserClip
	file, err := c.FormFile(reqeust.ClipFileName)
	if errors.Is(err, http.ErrNotMultipart) || errors.Is(err, http.ErrMissingFile) {
		var addUCR reqeust.AddUserClipRequest
		if err := c.ShouldBindJSON(&addUCR); err != nil {
			// 解析失败
			ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
			return
		}
		// 保存一个文本类型的记录到剪切板数据库中
		userClip = &model.UserClip{
			UserId:  userId,
			Type:    model.ClipText,
			Content: addUCR.Content,
			UseTime: time.Now(),
		}
	} else if err != nil {
		// 文件解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.FileParseError, err.Error())
		return
	} else {
		// 将上传的文件保存到文件路径中
		newNameFileName := util.GenerateNewNameForFile(file.Filename)
		filaSavePath := filepath.Join(config.ImageDirPath, newNameFileName)
		if err := c.SaveUploadedFile(file, filaSavePath); err != nil {
			ResponseFail(c, http.StatusInternalServerError, MyErr.FileSaveError, err.Error())
			return
		}
		// 将数据保存为一个图片类型记录
		userClip = &model.UserClip{
			UserId:  userId,
			Type:    model.ClipImage,
			Content: newNameFileName,
			UseTime: time.Now(),
		}
	}
	// 检查是否达到最大容量
	count, err := service.CountClipForUser(userId)
	if err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseQueryError, err.Error())
		return
	}
	if count >= config.UserClipMaxCapacity {
		// 删除使用时间最早的记录
		if err := service.DeleteUnusedClipForUser(userId); err != nil {
			ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseDeleteError, err.Error())
			return
		}
	}
	// 新记录落库
	if err := config.DataBase.Create(userClip).Error; err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseSaveError, "")
		return
	}
	// 成功响应
	ResponseSuccess(c)
}

func PageListUserClips(c *gin.Context) {
	// 拿到用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析分页请求
	var pageCondition *reqeust.PageDataRequest
	if err := c.ShouldBindJSON(&pageCondition); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	// 默认判断一下
	pageCondition = reqeust.JudgeAndSetDefaultPageDataRequest(pageCondition)
	// 进行查询
	pageUserClips, err := service.PageUserClips(userId, pageCondition.PageNum, pageCondition.PageSize)
	if err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseQueryError, err.Error())
		return
	}
	ResponseSuccessWithData(c, pageUserClips)
}

func DeleteUserClip(c *gin.Context) {
	// 拿到用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析删除请求
	var deleteRequest *reqeust.DeleteDataRequest
	if err := c.ShouldBindJSON(&deleteRequest); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	if deleteRequest.Id == nil {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataCannotEmpty, "")
		return
	}
	// 进行删除
	if err := service.DeleteUserClip(userId, *deleteRequest.Id); err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseDeleteError, err.Error())
		return
	}
}

func UpdateUserClipContent(c *gin.Context) {
	// 获取用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析更新请求
	var updateRequest *reqeust.UpdateUserClipRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	if updateRequest.Id == nil {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataCannotEmpty, "")
		return
	}
	// 这里只允许修改文本类型的剪切板内容，且只能修改为文本
	if err := service.UpdateUserClip(userId, *updateRequest.Id, updateRequest.Content); err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseUpdateError, err.Error())
		return
	}
	ResponseSuccess(c)
}

func UpdateUserClipUseTime(c *gin.Context) {
	// 获取用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析请求
	var updateUTR *reqeust.UpdateUserClipUseTimeRequest
	if err := c.ShouldBindJSON(&updateUTR); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	if updateUTR.Id == nil {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataCannotEmpty, "")
		return
	}
	// 更新使用时间
	if err := service.UpdateUserClipUseTime(userId, *updateUTR.Id); err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseUpdateError, err.Error())
		return
	}
	ResponseSuccess(c)
}
