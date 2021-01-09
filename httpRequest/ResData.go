package httpRequest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ResData struct {
	StatusCode int
	Data       *[]byte
}

/** 请求返回 */
func ResDataBack(resp *http.Response, data *[]byte) *ResData {

	byteArr := []string{}
	for _, v := range *data {
		byteArr = append(byteArr, strconv.FormatUint(uint64(v), 10))
	}
	fmt.Printf("[%s]\n", strings.Join(byteArr, ","))

	return &ResData{
		StatusCode: resp.StatusCode,
		Data:       data,
	}
}
