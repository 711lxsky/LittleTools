package view

import "time"

type TodoSimpleView struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
}

type TodoView struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	RemindAt time.Time `json:"remindAt"`
	Status   int       `json:"status"`
}
