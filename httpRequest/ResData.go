package httpRequest

import (
	"net/http"
)

type ResData struct {
	StatusCode int
	Data       *[]byte
}

/** 请求返回 */
func ResDataBack(resp *http.Response, data *[]byte) *ResData {
	return &ResData{
		StatusCode: resp.StatusCode,
		Data:       data,
	}
}
