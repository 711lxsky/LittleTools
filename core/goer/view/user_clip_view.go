package view

import "time"

type UserClipView struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	Type       int       `json:"type"`
	CreateTime time.Time `json:"createTime"`
}
