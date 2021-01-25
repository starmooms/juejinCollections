package collectReq

import (
	"io/ioutil"
	"juejinCollections/httpRequest"
	"juejinCollections/tool"

	"github.com/buger/jsonparser"
)

type Mock struct {
	Article     *[]byte
	Tags        *[]byte
	CollectData *[]byte
	Img         *[]byte
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
	collectData, _, _, err := jsonparser.Get(t, "collectData")
	tool.PanicErr(err)

	return &Mock{
		Tags:        &tags,
		Article:     &article,
		CollectData: &collectData,
	}
}

type MockReq struct {
	mock *Mock
}

func (m *MockReq) MockRequest(h *httpRequest.HttpRequest, next func() error) error {
	var mockData *[]byte
	switch h.Url {
	case GET_TAGSLIST:
		mockData = m.mock.Tags
		case GET_ARTICLE:
			mockData = m.mock.Article
		case GET_COLLECTDATA:
			mockData = m.mock.CollectData
		default:
			mockData = &[]byte{}
	}
	if mockData != nil {
		h.DoMock = func() (*[]byte, error) {
			return mockData, nil
		}
	}
	return next()
}

/** 生成mock请求方法 */
func GetMockRequest() httpRequest.MiddleFunc {
	m := &MockReq{
		mock: NewMock(),
	}
	return m.MockRequest
}
