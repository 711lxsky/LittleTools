package model

import (
	"gorm.io/gorm"
	"time"
)

// UserClip 用户关联剪切板的实体映射
type UserClip struct {
	gorm.Model
	ID      int       `json:"id" gorm:"primaryKey;not null"`                // 主键id
	UserId  int       `json:"userId" gorm:"type:int;not null;default:0"`    // 关联的用户id
	Type    int       `json:"type" gorm:"type:int;not null;default:1"`      // 存储的数据类型
	Content string    `json:"content" gorm:"type:text;not null;default:''"` // 存储的数据内容
	UseTime time.Time `gorm:"type:text;default:CURRENT_TIMESTAMP"`          // 使用时间
}
