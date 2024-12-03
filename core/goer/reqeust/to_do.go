package reqeust

import "time"

type AddTodoRequest struct {
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	RemindAt *time.Time `json:"remindAt,omitempty"`
}

type UpdateTodoRequest struct {
	Id       *int       `json:"id"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	RemindAt *time.Time `json:"remindAt,omitempty"`
	Status   *int       `json:"status"`
}
