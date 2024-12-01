package model

import (
	"github.com/jinzhu/gorm"
)

// User 表users的实体映射
type User struct {
	gorm.Model
	ID       int    `json:"userId" gorm:"primaryKey;not null"`
	UserName string `json:"userName" gorm:"type:text;not null;default:''"`
	Password string `json:"password" gorm:"type:text;not null;default:''"`
}
