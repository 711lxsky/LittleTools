package service

import (
	"github.com/robfig/cron/v3"
	"goer/config"
	"goer/model"
	"goer/util"
	"goer/view"
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

func PageTodos(userId, pageNum, pageSize int) (*view.PageDataView, error) {
	// 先拿到总记录数
	var total int64
	if err := config.DataBase.Model(&model.Todo{}).Where("user_id = ?", userId).Count(&total).Error; err != nil {
		return nil, err
	}
	totalInt := int(total)
	// 进行查询
	offset := util.Max((pageNum-1)*pageSize, totalInt)
	var todos []model.Todo
	if err := config.DataBase.Model(&model.Todo{}).Where("user_id = ?", userId).Order("created_at desc").Offset(offset).Limit(pageSize).Find(&todos).Error; err != nil {
		return nil, err
	}
	// 封装数据
	var data []interface{}
	for _, todo := range todos {
		// 将字段属性直接转换给视图结构
		tdv := &view.TodoSimpleView{
			Id:     todo.ID,
			Title:  todo.Title,
			Status: todo.Status,
		}
		data = append(data, tdv)
	}
	return view.NewPageDataView(data, pageNum, pageSize, totalInt), nil
}

func GetTodoInfo(userId, todoId int) (*view.TodoView, error) {
	var todo model.Todo
	if err := config.DataBase.Where("id = ? and user_id = ?", todoId, userId).First(&todo).Error; err != nil {
		return nil, err
	}
	// 转换为视图返回
	tdv := &view.TodoView{
		Id:       todo.ID,
		Title:    todo.Title,
		Content:  todo.Content,
		RemindAt: todo.RemindAt,
		Status:   todo.Status,
	}
	return tdv, nil
}

func DeleteTodo(userId, todoId int) error {
	return config.DataBase.Where("id = ? and user_id = ?", todoId, userId).Delete(&model.Todo{}).Error
}

func UpdateTodo(userId, todoId int, Title, Content string, remindAt *time.Time, status *int) error {
	var todo model.Todo
	config.DataBase.Where("id = ? and user_id = ?", todoId, userId).First(&todo)
	if "" != Title {
		todo.Title = Title
	}
	if "" != Content {
		todo.Content = Content
	}
	if remindAt != nil {
		todo.RemindAt = *remindAt
	}
	if status != nil {
		todo.Status = *status
	}
	return config.DataBase.Save(&todo).Error
}

func TimedRemindTodoWithEmail() {
	log.Printf("定时提醒待办事项")
	c := cron.New()
	if _, err := c.AddFunc("@every 1m", func() {
		var todos []model.Todo
		now := time.Now()
		if err := config.DataBase.Where("status = ? and remind_at < ?", model.StatusWaitingRemind, now).Find(&todos).Error; err != nil {
			log.Panicf("failed to find expired todos: %v", err)
		}
		for _, todo := range todos {
			go func(todo model.Todo) {
				sendRemindEmail(todo)
				if err := config.DataBase.Where("id = ?", todo.ID).Update("status", model.StatusReminded).Error; err != nil {
					log.Panicf("failed to update todo status: %v", err)
				}
			}(todo)
		}
	}); err != nil {
		log.Panicf("failed to add cron job: %v", err)
	}
	c.Start()
	// 防止主程序退出
	select {}
}

func sendRemindEmail(todo model.Todo) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailSendUser)
	email := GetEmailByUserId(todo.UserId)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "待办事项提醒"+todo.Title)
	m.SetBody("text/plain", "Hi, "+email+"，"+todo.Content)
	d := util.BuildDialer()
	if err := d.DialAndSend(m); err != nil {
		log.Printf("failed to send email: %v", err)
	}
}
