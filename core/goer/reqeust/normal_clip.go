package reqeust

type NormalClipboardRequest struct {
	Content string `json:"text"`
}

var (
	NormalClipFileName = "pic"
)
