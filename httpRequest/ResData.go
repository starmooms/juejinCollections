package httpRequest

import (
	"net/http"
)

type ResData struct {
	Resp       *http.Response
	StatusCode int
	Data       *[]byte
}

/** 请求返回 */
func ResDataBack(resp *http.Response, data *[]byte) *ResData {

	// byteArr := []string{}
	// for _, v := range *data {
	// 	byteArr = append(byteArr, strconv.FormatUint(uint64(v), 10))
	// }
	// fmt.Printf("[%s]\n", strings.Join(byteArr, ","))

	return &ResData{
		Resp:       resp,
		StatusCode: resp.StatusCode,
		Data:       data,
	}
}
