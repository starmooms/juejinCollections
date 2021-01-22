package collectReq

import (
	"io/ioutil"
	"juejinCollections/httpRequest"
	"juejinCollections/tool"
	"net/http"

	"github.com/buger/jsonparser"
)

type Mock struct {
	Article *[]byte
	Tags    *[]byte
}

/** 创建mock数据 */
func NewMock() *Mock {
	var err error
	t, err := ioutil.ReadFile("./collectReq/mock.json")
	tool.PanicErr(err)

	tags, _, _, err := jsonparser.Get(t, "tags")
	tool.PanicErr(err)
	article, _, _, err := jsonparser.Get(t, "article")
	tool.PanicErr(err)

	return &Mock{
		Tags:    &tags,
		Article: &article,
	}
}

type MockReq struct {
	mock *Mock
}

func (m *MockReq) MockRequest(h *httpRequest.HttpRequest) *httpRequest.HttpRequest {
	var mockData *[]byte
	switch h.Url {
	case GET_TAGSLIST:
		mockData = m.mock.Tags
	case GET_ARTICLE:
		mockData = m.mock.Article
	}
	h.DoMock = func() (*httpRequest.ResData, error) {
		resp := &http.Response{}
		return httpRequest.ResDataBack(resp, mockData), nil
	}
	return httpRequest.Request(h)
}

/** 生成mock请求方法 */
func GetMockRequest() func(h *httpRequest.HttpRequest) *httpRequest.HttpRequest {
	m := &MockReq{
		mock: NewMock(),
	}
	return m.MockRequest
}
