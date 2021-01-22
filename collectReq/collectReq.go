package collectReq

import (
	"encoding/json"
	"juejinCollections/dal"
	"juejinCollections/httpRequest"
	"juejinCollections/logger"
	"juejinCollections/model"
	"juejinCollections/tool"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
)

var log = logger.GetLog()
var request = httpRequest.Request
var userMock = true

func init() {
	if userMock {
		request = GetMockRequest()
	}
}

func Run() {
	wg := tool.NewWaitGroup(10)
	wg.Add(func() {
		wg.Add(func() {
			GetTagList(1116759544852221)
		})
		wg.Add(func() {
			GetArticle("6844904034181070861")
		})
	})
	wg.Wait()
}

// 获取收藏列表
func GetTagList(userId int) (err error) {
	httpReq := request(&httpRequest.HttpRequest{
		Url:    GET_TAGSLIST,
		Method: "GET",
		Params: &gin.H{
			"user_id": userId,
			"cursor":  0,
			"limit":   200,
		},
	})
	result, err := httpReq.DoRequest()
	if err != nil {
		tool.BackError(err)
		return err
	}
	reqCollectList := &CollectListStruct{}
	json.Unmarshal(*result.Data, reqCollectList)

	if _, err := dal.AddTags(&reqCollectList.Data); err != nil {
		return tool.BackError(err)
	}

	return nil
}

// 通过id获取文章
func GetArticle(id string) error {
	httpReq := request(&httpRequest.HttpRequest{
		Url:    GET_ARTICLE,
		Method: "POST",
		Params: &gin.H{
			"article_id": id,
		},
	})
	var err error
	result, err := httpReq.DoRequest()
	if err != nil {
		return tool.BackError(err)
	}

	res := &ArticleRes{}
	err = json.Unmarshal(*result.Data, res)
	if err != nil {
		return tool.BackError(err)
	}
	err = CheckErr(res)

	articleByt, _, _, err := jsonparser.Get(*result.Data, "data", "article_info")
	if err != nil {
		return tool.BackError(err)
	}

	article := &model.ArticleModel{}
	if err = json.Unmarshal(articleByt, article); err != nil {
		return tool.BackError(err)
	}

	_, err = dal.AddArticle(article)
	if err != nil {
		return tool.BackError(err)
	}
	return nil
}

// 获取收藏内容
func GetTagContext() {

}
