package reqeust

type ImageNormalClipRequest struct {
	Password string `form:"password"`
}

type TextNormalClipRequest struct {
	Content  string  `json:"text"`
	Password *string `json:"password"`
}

type GetNormalClipRequest struct {
	Password string `json:"password"`
}

var (
	ClipFileName = "pic"
)
