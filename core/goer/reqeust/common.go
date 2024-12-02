package reqeust

import "goer/config"

type PageDataRequest struct {
	PageNum  int
	PageSize int
}

func JudgeAndSetDefaultPageDataRequest(pageDataRequest *PageDataRequest) *PageDataRequest {
	if pageDataRequest.PageNum <= 0 {
		pageDataRequest.PageNum = config.PageDataDefaultNum
	}
	if pageDataRequest.PageSize <= 0 {
		pageDataRequest.PageSize = config.PageDataDefaultSize
	}
	return pageDataRequest
}

type DeleteDataRequest struct {
	Id int `json:"id"`
}
