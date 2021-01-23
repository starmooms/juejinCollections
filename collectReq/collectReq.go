package collectReq

import (
	"encoding/json"
	"juejinCollections/dal"
	"juejinCollections/httpRequest"
	"juejinCollections/logger"
	"juejinCollections/model"
	"juejinCollections/tool"

	"github.com/buger/jsonparser"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

var log = logger.GetLog()
var requestWrap = &httpRequest.RequestWarp{}
var request = requestWrap.GetNewRequest
var userMock = true

func init() {
	if userMock {
		requestWrap.Use(GetMockRequest())
	}

	// 检查请求结果
	requestWrap.Use(func(h *httpRequest.HttpRequest, next func() error) error {
		if err := next(); err != nil {
			return err
		}

		rb := &ResBase{}
		if err := json.Unmarshal(*h.ResData.Data, rb); err != nil {
			return errors.Wrap(err, "checkRes Unmarshal Err ")
		}

		if err := rb.CheckErr(); err != nil {
			return err
		}

		return nil
	})

}

func Run() {
	wg := tool.NewWaitGroup(10)
	wg.Add(func() {
		GetTagList("1116759544852221")
	})
	// wg.Add(func() {
	// 	GetArticle("6844904034181070861")
	// })
	// wg.Add(func() {
	// 	err := GetCollectData("6889585424071655431", 10)
	// 	if err != nil {
	// 		tool.BackError(err)
	// 	}
	// })
	wg.Wait()
}

// 获取收藏列表
func GetTagList(userId string) (err error) {
	reqCollectList := &CollectListStruct{}
	httpReq, err := request(&httpRequest.HttpRequest{
		Url:    GET_TAGSLIST,
		Method: "GET",
		Params: &gin.H{
			"user_id": userId,
			"cursor":  0,
			"limit":   200,
		},
		ResJson: reqCollectList,
	})
	if err != nil {
		return tool.BackError(err)
	}

	_, err = httpReq.DoRequest()
	if err != nil {
		return tool.BackError(err)
	}
	// reqCollectList := &CollectListStruct{}
	// json.Unmarshal(*result.Data, reqCollectList)

	if _, err := dal.AddTags(&reqCollectList.Data); err != nil {
		return tool.BackError(err)
	}

	return nil
}

// 通过id获取文章
func GetArticle(id string) (err error) {
	httpReq, err := request(&httpRequest.HttpRequest{
		Url:    GET_ARTICLE,
		Method: "POST",
		Params: &gin.H{
			"article_id": id,
		},
	})
	if err != nil {
		return tool.BackError(err)
	}

	result, err := httpReq.DoRequest()
	if err != nil {
		return tool.BackError(err)
	}

	res := &ArticleRes{}
	err = json.Unmarshal(*result.Data, res)
	if err != nil {
		return tool.BackError(err)
	}

	articleByt, _, _, err := jsonparser.Get(*result.Data, "data", "article_info")
	if err != nil {
		return tool.BackError(err)
	}

	article := &model.ArticleModel{}
	if err = json.Unmarshal(articleByt, article); err != nil {
		return tool.BackError(err)
	}

	_, err = dal.AddArticle(&[]*model.ArticleModel{article})
	if err != nil {
		return tool.BackError(err)
	}
	return nil
}

// 获取收藏内容
func GetCollectData(tagId string, cursor int) (err error) {
	httpReq, err := request(&httpRequest.HttpRequest{
		Url:    GET_COLLECTDATA,
		Method: "POST",
		Params: &gin.H{
			"tag_id": tagId,
			"cursor": cursor,
		},
	})
	if err != nil {
		return err
	}

	result, err := httpReq.DoRequest()
	if err != nil {
		return err
	}

	collectData := []*model.ArticleModel{}
	tagArticle := []*model.TagArticleModel{}
	var jsonErr *error = nil
	_, err = jsonparser.ArrayEach(*result.Data, func(value []byte, dataType jsonparser.ValueType, offset int, eachErr error) {
		if jsonErr != nil {
			return
		}
		artByt, _, _, err := jsonparser.Get(value, "article_info")
		if err != nil {
			err = errors.Wrap(err, "ArrayEach Get 'article_info' Error")
			jsonErr = &err
			return
		}

		artItem := &model.ArticleModel{}
		err = json.Unmarshal(artByt, artItem)
		if err != nil {
			err = errors.Wrap(err, "ArrayEach Get 'article Unmarshal' Error")
			jsonErr = &err
			return
		}
		collectData = append(collectData, artItem)
		tagArticle = append(tagArticle, &model.TagArticleModel{
			TagId:     tagId,
			ArticleId: artItem.ArticleId,
		})

	}, "data", "article_list")

	if err == nil && jsonErr != nil {
		err = *jsonErr
	}
	if err != nil {
		return err
	}

	if _, err = dal.AddArticle(&collectData); err != nil {
		return err
	}

	if _, err = dal.AddTagArticle(&tagArticle); err != nil {
		return err
	}

	return nil
}
