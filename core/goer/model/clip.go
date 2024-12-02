package model

import (
	"crypto/rand"
	"encoding/base64"
	"gorm.io/gorm"
)

type Clip struct {
	gorm.Model
	ID         int    `json:"id" gorm:"primaryKey;not null"`                   // 主键id
	Type       int    `json:"type" gorm:"type:int;not null;default:1"`         // 存储的数据类型
	Content    string `json:"content" gorm:"type:text;not null;default:''"`    // 存储的数据内容
	Identifier string `json:"identifier" gorm:"type:text;not null;default:''"` // 剪切内容对应的标记
	Password   string `json:"password" gorm:"type:text;not null;default:''"`
}

var (
	ClipText  = 1
	ClipImage = 2
)

// GenerateRandomString 生成一个指定长度的随机字符串
func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
