package service

import (
	"goer/config"
	"goer/model"
	"strings"
)

func GetUserByName(userName string) (user *model.User, err error) {
	user = &model.User{}
	err = config.DataBase.Where("user_name = ?", userName).First(user).Error
	return
}

func CheckUserEmailExist(userId int) bool {
	var userFromDB model.User
	if err := config.DataBase.Where("id = ?", userId).First(&userFromDB).Error; err != nil {
		return false
	}
	return strings.TrimSpace(userFromDB.Email) != ""
}

func GetEmailByUserId(userId int) string {
	var userFromDB model.User
	if err := config.DataBase.Where("id = ?", userId).First(&userFromDB).Error; err != nil {
		return ""
	}
	return userFromDB.Email
}

func UpdateUserInfo(userId int, userName, email string) error {
	var userInDB model.User
	if err := config.DataBase.Where("id = ?", userId).First(&userInDB).Error; err != nil {
		return err
	}
	if "" != userName {
		userInDB.UserName = userName
	}
	if "" != email {
		userInDB.Email = email
	}
	return config.DataBase.Save(&userInDB).Error
}
