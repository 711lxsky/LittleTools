package reqeust

type AddUserClipRequest struct {
	Content string `json:"text"`
}

type UpdateUserClipRequest struct {
	Id      *int   `json:"id"`
	Content string `json:"content"`
}

type UpdateUserClipUseTimeRequest struct {
	Id *int `json:"id"`
}
