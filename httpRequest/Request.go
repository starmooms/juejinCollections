package httpRequest

import "strings"

/** 初始化请求 */
func Request(h *HttpRequest) *HttpRequest {
	h.Method = strings.ToUpper(h.Method)
	h.NewRequest()
	return h
}
