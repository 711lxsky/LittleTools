package reqeust

import "goer/config"

type PageDataRequest struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
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
	Id *int `json:"id"`
}

type ViewDetailRequest struct {
	Id *int `json:"id"`
}
