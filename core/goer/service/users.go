package service

import (
	"goer/config"
	"goer/model"
)

func GetUserByName(userName string) (user *model.User, err error) {
	user = &model.User{}
	err = config.DataBase.Where("user_name = ?", userName).First(user).Error
	return
}
