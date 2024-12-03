package model

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	ID       int       `json:"id" gorm:"primaryKey;not null"`
	UserId   int       `json:"userId" gorm:"type:int;not null;default:0"` // 关联的用户id
	Title    string    `json:"title" gorm:"type:text;not null;default:''"`
	Content  string    `json:"content" gorm:"type:text;not null;default:''"`
	RemindAt time.Time `json:"remindAt" gorm:"type:text;default:CURRENT_TIMESTAMP"`
	Status   int       `json:"status" gorm:"type:int;not null;default:0"`
}

var (
	StatusAdd           = 1
	StatusDone          = 2
	StatusWaitingRemind = 3
	StatusReminded      = 4
)
