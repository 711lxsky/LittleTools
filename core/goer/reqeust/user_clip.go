package reqeust

type AddUserClipRequest struct {
	Content string `json:"content"`
}

type UpdateUserClipRequest struct {
	Id      *int   `json:"id"`
	Content string `json:"content"`
}

type UpdateUserClipUseTimeRequest struct {
	Id *int `json:"id"`
}
