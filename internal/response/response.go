package response

type Response struct {
	ErrorCode int         `json:"error_code"`
	Message   string      `json:"message"`
	Count     int         `json:"count,omitempty"`
	Page      int         `json:"page,omitempty"`
	Size      int         `json:"size,omitempty"`
	Total     interface{} `json:"total,omitempty"`
	Today     interface{} `json:"today,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type SuccessOptions struct {
	Count int
	Page  int
	Size  int
}

type SuccessOption func(req *SuccessOptions)

func WithPagination(count, page, size int) SuccessOption {
	return func(req *SuccessOptions) {
		req.Count = count
		req.Page = page
		req.Size = size
	}
}

func SuccessWithTotal(data, total, today interface{}, opts ...SuccessOption) interface{} {
	res := &SuccessOptions{}
	for _, opt := range opts {
		opt(res)
	}
	return Response{
		ErrorCode: 0,
		Message:   "Success",
		Page:      res.Page,
		Count:     res.Count,
		Size:      res.Size,
		Total:     total,
		Today:     today,
		Data:      data,
	}
}

func Success(data interface{}, opts ...SuccessOption) interface{} {
	res := &SuccessOptions{}
	for _, opt := range opts {
		opt(res)
	}
	return Response{
		ErrorCode: 0,
		Message:   "Success",
		Page:      res.Page,
		Count:     res.Count,
		Size:      res.Size,
		Data:      data,
	}
}

func ERROR(err error) interface{} {
	return Response{
		Message:   err.Error(),
		ErrorCode: 1,
	}
}
