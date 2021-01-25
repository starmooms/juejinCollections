package httpRequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"juejinCollections/logger"
	"juejinCollections/tool"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

// func Request(method string, url string, params *gin.H) *HttpRequest {
// 	method = strings.ToUpper(method)
// 	return &HttpRequest{
// 		method: method,
// 		url:    url,
// 		params: params,
// 	}
// }

var log = logger.Logger

type HttpRequest struct {
	Url     string
	Method  string
	Params  *gin.H
	Req     *http.Request
	ResData *ResData
	ResJson interface{}
	DoMock  func() (*[]byte, error)
	rw      *RequestWarp
}

func (h *HttpRequest) setQuery() {
	if h.Params == nil {
		return
	}

	req := h.Req
	q := req.URL.Query()
	params := *h.Params

	for key, val := range params {
		value := ""
		vType := reflect.TypeOf(val).Kind()
		switch vType {
		case reflect.String:
			value = val.(string)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(int64(val.(int)), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = strconv.FormatUint(uint64(val.(uint64)), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(float64(val.(float64)), 'f', -1, 64)
		case reflect.Bool:
			value = strconv.FormatBool(val.(bool))
		}

		q.Add(key, value)
		// q.Add(p, strings(val))
	}
	req.URL.RawQuery = q.Encode()
}

func (h *HttpRequest) setJSONBody() error {
	if h.Params != nil {
		params := *h.Params
		byts, err := json.Marshal(params)
		if err != nil {
			return errors.Wrap(err, "setJSONBody Marshal Params Error")
		}
		h.Req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		h.Req.ContentLength = int64(len(byts))
		h.SetHeader("Content-Type", "application/json")
	}
	return nil
}

func (h *HttpRequest) SetHeader(k, v string) {
	h.Req.Header.Set(k, v)
}

func (h *HttpRequest) NewRequest() (*http.Request, error) {
	h.Method = strings.ToUpper(h.Method)
	req, err := http.NewRequest(h.Method, h.Url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "NewRequest Error")
	}
	h.Req = req
	// h.SetHeader("Content-Type", "application/json")
	h.SetHeader("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	return req, nil
}

func (h *HttpRequest) DoRequest() (data *ResData, err error) {
	next := func() error {
		data, err = h.Do()
		h.ResData = data
		if h.ResJson != nil {
			err = json.Unmarshal(*data.Data, h.ResJson)
			if err != nil {
				return errors.Wrap(err, "DoRequest Unmarshal Error")
			}
		}
		return err
	}

	middleList := h.rw.middle
	middleLen := len(middleList)

	if middleLen > 0 {
		for i := middleLen - 1; i >= 0; i-- {
			middleF := h.rw.middle[i]
			lastNext := next
			next = func() error {
				return middleF(h, lastNext)
			}
		}
	}

	err = next()
	if err != nil {
		h.PrintReq(true)
		return nil, err
	}

	h.PrintReq(false)
	return data, err
}

func (h *HttpRequest) Do() (*ResData, error) {
	switch h.Method {
	case "GET":
		h.setQuery()
	case "POST":
		err := h.setJSONBody()
		if err != nil {
			return nil, err
		}
	}

	if h.DoMock != nil {
		resp := &http.Response{}
		mockData, err := h.DoMock()
		if err != nil {
			return nil, err
		}
		return ResDataBack(resp, mockData), nil
	}

	resp, err := http.DefaultClient.Do(h.Req)
	if err != nil {
		return nil, errors.Wrap(err, "HttpRequest Do Error")
	}
	defer resp.Body.Close()

	// if resp.Body == nil {
	// 	return nil, nil
	// }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "HttpRequest Read Body Error")
	}

	return ResDataBack(resp, &body), nil
}

func (h *HttpRequest) PrintReq(isErr bool) {
	var err error
	reqBody := "null"
	if h.Req.Body != nil {
		reqBodyByt, err := ioutil.ReadAll(h.Req.Body)
		if err != nil {
			tool.BackError(errors.Wrap(err, "PrintReq Read Body Err"))
		}
		reqBody = string(reqBodyByt)
	}

	query := h.Req.URL.Query()
	reqQuery := map[string]interface{}{}
	for key, item := range query {
		var v interface{}
		if len(item) == 1 {
			v = item[0]
		} else {
			v = item
		}
		reqQuery[key] = v
	}

	params := gin.H{}
	if h.Params != nil {
		params = *h.Params
	}

	statucCode := h.ResData.StatusCode

	data := map[string]interface{}{
		"url":    h.Url,
		"params": params,
		"request": map[string]interface{}{
			"Content-Type": h.Req.Header.Get("Content-Type"),
			"body":         string(reqBody),
			"query":        reqQuery,
		},
		"respond": map[string]interface{}{
			"Status Code": h.ResData.StatusCode,
			"data":        tool.LimtStr(string(*h.ResData.Data), 200),
		},
	}
	dataByt, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		tool.BackError(errors.Wrap(err, "PrintReq MarshalIndent Err"))
	}

	dataStr := string(dataByt)

	if isErr || (statucCode != 200 && statucCode != 0) {
		log.Error(dataStr)
	} else {
		log.Info(dataStr)
	}
}
