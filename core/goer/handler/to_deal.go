package handler

import (
	"github.com/gin-gonic/gin"
	"goer/config"
	MyErr "goer/error"
	"goer/model"
	"goer/reqeust"
	"goer/service"
	"net/http"
	"strings"
	"time"
)

func AddTodo(c *gin.Context) {
	// 拿到用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析新增待办请求
	var addTodoR *reqeust.AddTodoRequest
	if err := c.ShouldBindJSON(&addTodoR); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	// 检查标题是否为空
	if addTodoR.Title == "" {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataCannotEmpty, "")
		return
	}
	// 查看是否有提醒时间
	status := model.StatusAdd
	if addTodoR.RemindAt != nil {
		// 检查邮箱是否绑定
		if !service.CheckUserEmailExist(userId) {
			ResponseFail(c, http.StatusBadRequest, MyErr.UserEmailError, "")
		}
		status = model.StatusWaitingRemind
	}
	// 新建一条待办
	var todo *model.Todo
	todo = &model.Todo{
		Title:    addTodoR.Title,
		Content:  addTodoR.Content,
		UserId:   userId,
		RemindAt: *addTodoR.RemindAt,
		Status:   status,
	}
	if err := config.DataBase.Create(todo).Error; err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseSaveError, "")
		return
	}
	// 成功
	ResponseSuccess(c)
}

func PageTodoList(c *gin.Context) {
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
	// 服务层查询
	pageTodos, err := service.PageTodos(userId, pageCondition.PageNum, pageCondition.PageSize)
	if err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseQueryError, err.Error())
	}
	ResponseSuccessWithData(c, pageTodos)
}

func GetTodoDetailInfo(c *gin.Context) {
	// 拿到用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析查询详情请求
	var detailR *reqeust.ViewDetailRequest
	if err := c.ShouldBindJSON(&detailR); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	if detailR.Id == nil {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataCannotEmpty, "")
	}
	// 下到服务层进行查询
	todoInfo, err := service.GetTodoInfo(userId, *detailR.Id)
	if err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseQueryError, err.Error())
	}
	ResponseSuccessWithData(c, todoInfo)
}

func DeleteTodo(c *gin.Context) {
	// 拿到用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析删除待办请求
	var deleteR *reqeust.DeleteDataRequest
	if err := c.ShouldBindJSON(&deleteR); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	if deleteR.Id == nil {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataCannotEmpty, "")
	}
	// 下到服务层进行删除
	if err := service.DeleteTodo(userId, *deleteR.Id); err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseDeleteError, err.Error())
	}
	ResponseSuccess(c)
}

func UpdateTodo(c *gin.Context) {
	// 拿到用户id
	userId := getUserIdFromContext(c)
	if userId == MyErr.IntErrValue {
		return
	}
	// 解析更新请求
	var updateR *reqeust.UpdateTodoRequest
	if err := c.ShouldBindJSON(&updateR); err != nil {
		// 解析失败
		ResponseFail(c, http.StatusBadRequest, MyErr.JSONParseError, err.Error())
		return
	}
	if updateR.Id == nil {
		ResponseFail(c, http.StatusBadRequest, MyErr.DataCannotEmpty, "")
	}
	remindTime := updateR.RemindAt
	updateStatus := updateR.Status
	if remindTime != nil {
		// 检查时间是否合法
		if remindTime.Before(time.Now()) {
			ResponseFail(c, http.StatusBadRequest, MyErr.DataLogicError, "")
		}
		// 检查邮箱是否绑定
		if !service.CheckUserEmailExist(userId) {
			ResponseFail(c, http.StatusBadRequest, MyErr.UserEmailError, "")
		}
		// 检查在此情况下状态是否合法
		if updateStatus != nil && *updateStatus == model.StatusDone {
			ResponseFail(c, http.StatusBadRequest, MyErr.DataLogicError, "")
		}
	} else {
		if updateStatus != nil && *updateStatus == model.StatusWaitingRemind {
			ResponseFail(c, http.StatusBadRequest, MyErr.DataLogicError, "")
		}
	}
	if err := service.UpdateTodo(userId, *updateR.Id, strings.TrimSpace(updateR.Title), strings.TrimSpace(updateR.Content), remindTime, updateStatus); err != nil {
		ResponseFail(c, http.StatusInternalServerError, MyErr.DataBaseUpdateError, err.Error())
	}
	ResponseSuccess(c)
}
