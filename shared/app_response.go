package shared

type AppResponse struct {
	Data   any `json:"data"`
	Filter any `json:"filter,omitempty"`
	Paging any `json:"paging,omitempty"`
}

func NewAppResponse(data any, filter any, paging any) *AppResponse {
	return &AppResponse{
		Data:   data,
		Filter: filter,
		Paging: paging,
	}
}

func SimpleResponse(data any) *AppResponse {
	return NewAppResponse(data, nil, nil)
}