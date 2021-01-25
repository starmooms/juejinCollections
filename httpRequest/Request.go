package httpRequest

type MiddleFunc = func(h *HttpRequest, next func() error) error

type RequestWarp struct {
	middle []MiddleFunc
}

func (r *RequestWarp) Use(n ...MiddleFunc) {
	r.middle = append(r.middle, n...)
}

func (r *RequestWarp) GetNewRequest(h *HttpRequest) (*HttpRequest, error) {
	h.rw = r
	_, err := h.NewRequest()
	return h, err
}

/** 初始化请求 */
func Request(h *HttpRequest) *HttpRequest {
	h.NewRequest()
	return h
}
