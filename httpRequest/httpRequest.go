package httpRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

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

type HttpRequest struct {
	Url    string
	Method string
	Params *gin.H
	Req    *http.Request
	DoMock func() (*ResData, error)
}

func (h *HttpRequest) setQuery() {
	req := h.Req
	q := req.URL.Query()
	params := *h.Params

	for key, val := range params {
		value := ""
		// switch val.(type) {
		// case string:
		// 	value = val.(string)
		// case int, int8, int16, int32, int64:
		// 	value = strconv.Itoa(val.(int64))
		// case uint, uint8, uint16, uint32, uint64:
		// 	value = strconv.FormatUint(val.(uint64), 10)
		// case float32, float64:
		// 	value = strconv.FormatFloat(val.(float64), 'f', -1, 64)
		// case bool:
		// 	value = strconv.FormatBool(val.(bool))
		// }

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

		fmt.Println("key, value", key, value)
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
			return err
		}
		h.Req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		h.Req.ContentLength = int64(len(byts))
		h.SetHeader("Content-Type", "application/json")
		return nil
	}
	return nil
}

func (h *HttpRequest) SetHeader(k, v string) {
	h.Req.Header.Set(k, v)
}

func (h *HttpRequest) NewRequest() (*http.Request, error) {
	req, err := http.NewRequest(h.Method, h.Url, nil)
	if err != nil {
		return nil, err
	}
	h.Req = req
	// h.SetHeader("Content-Type", "application/json")
	h.SetHeader("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	return req, nil
}

func (h *HttpRequest) DoRequest() (*ResData, error) {
	if h.DoMock != nil {
		return h.DoMock()
	}

	switch h.Method {
	case "GET":
		h.setQuery()
	case "POST":
		err := h.setJSONBody()
		if err != nil {
			return nil, err
		}
	}

	resp, err := http.DefaultClient.Do(h.Req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// if resp.Body == nil {
	// 	return nil, nil
	// }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ResDataBack(resp, &body), nil
}

// func Get(url string, params *gin.H) {
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64
// 	q := req.URL.Query()
// 	for key, val := range *params {
// 		value := ""
// 		// switch val.(type) {
// 		// case string:
// 		// 	value = val.(string)
// 		// case int, int8, int16, int32, int64:
// 		// 	value = strconv.Itoa(val.(int64))
// 		// case uint, uint8, uint16, uint32, uint64:
// 		// 	value = strconv.FormatUint(val.(uint64), 10)
// 		// case float32, float64:
// 		// 	value = strconv.FormatFloat(val.(float64), 'f', -1, 64)
// 		// case bool:
// 		// 	value = strconv.FormatBool(val.(bool))
// 		// }

// 		vType := reflect.TypeOf(val).Kind()
// 		switch vType {
// 		case reflect.String:
// 			value = val.(string)
// 		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 			value = strconv.FormatInt(int64(val.(int)), 10)
// 		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 			value = strconv.FormatUint(uint64(val.(uint64)), 10)
// 		case reflect.Float32, reflect.Float64:
// 			value = strconv.FormatFloat(float64(val.(float64)), 'f', -1, 64)
// 		case reflect.Bool:
// 			value = strconv.FormatBool(val.(bool))
// 		}

// 		fmt.Println("key, value", key, value)
// 		q.Add(key, value)
// 		// q.Add(p, strings(val))
// 	}
// 	req.URL.RawQuery = q.Encode()

// 	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

// 	resp, err := http.DefaultClient.Do(req)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(string(body))
// }

// func Post(url string, params *gin.H) (data *gin.H, err error) {
// 	req, err := newRequest("POST", url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if params != nil {
// 		byts, err := json.Marshal(params)
// 		if err != nil {
// 			return nil, err
// 		}
// 		req.Body = ioutil.NopCloser(bytes.NewReader(byts))
// 		req.ContentLength = int64(len(byts))
// 		req.Header.Set("Content-Type", "application/json")
// 	}
// }
