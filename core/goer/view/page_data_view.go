package view

type PageDataView struct {
	Data     []interface{}
	Total    int
	PageSize int
	PageNum  int
}

func NewPageDataView(data []interface{}, total int, pageSize int, pageNum int) *PageDataView {
	return &PageDataView{
		Data:     data,
		Total:    total,
		PageSize: pageSize,
		PageNum:  pageNum,
	}
}
