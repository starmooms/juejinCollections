package collectReq

import (
	"encoding/json"
	"juejinCollections/dal"
	"juejinCollections/httpRequest"
	"juejinCollections/logger"
	"juejinCollections/model"
	"juejinCollections/tool"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

var log = logger.Logger

var requestWrap = &httpRequest.RequestWarp{}
var request = requestWrap.GetNewRequest

var imgRequestWrap = &httpRequest.RequestWarp{}
var imgRequest = imgRequestWrap.GetNewRequest

var userMock = false

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
	ac := NewAction("1116759544852221")
	// ac.DbArticleId = []string{"6844903480126078989"}
	ac.Run()
}

// 获取收藏列表
func GetTagList(userId string) (_ *[]model.Tag, err error) {
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
		return nil, err
	}

	_, err = httpReq.DoRequest()
	if err != nil {
		return nil, err
	}

	tagList := &reqCollectList.Data
	if _, err := dal.AddTags(tagList); err != nil {
		return nil, err
	}

	return tagList, nil
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

	article := &model.Article{}
	if err = json.Unmarshal(articleByt, article); err != nil {
		return tool.BackError(err)
	}

	_, err = dal.AddArticle(&[]*model.Article{article})
	if err != nil {
		return tool.BackError(err)
	}
	return nil
}

// 获取收藏内容
func GetCollectData(tagId string, cursor int) (collectData *CollectArticle, articleListPtr *[]*model.Article, err error) {
	collectData = &CollectArticle{}
	cursorStr := strconv.Itoa(cursor)

	httpReq, err := request(&httpRequest.HttpRequest{
		Url:    GET_COLLECTDATA,
		Method: "POST",
		Params: &gin.H{
			"collection_id": tagId,
			"cursor":        cursorStr,
			"limit":         10,
		},
		ResJson: collectData,
	})
	if err != nil {
		return
	}

	result, err := httpReq.DoRequest()
	if err != nil {
		return
	}

	articleList := []*model.Article{}
	tagArticle := []*model.TagArticleId{}
	var jsonErr *error = nil

	c, _, _, err := jsonparser.Get(*result.Data, "has_more")
	if err != nil {
		log.Info(err)
	} else {
		log.Info(c)
	}

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

		artItem := &model.Article{}
		err = json.Unmarshal(artByt, artItem)
		if err != nil {
			err = errors.Wrap(err, "ArrayEach Get 'article Unmarshal' Error")
			jsonErr = &err
			return
		}

		articleList = append(articleList, artItem)
		tagArticle = append(tagArticle, &model.TagArticleId{
			TagId:     tagId,
			ArticleId: artItem.ArticleId,
		})

	}, "data", "articles")

	if err == nil && jsonErr != nil {
		err = *jsonErr
	}
	if err != nil {
		return
	}

	articleListPtr = &articleList
	if _, err = dal.AddArticle(articleListPtr); err != nil {
		return
	}

	if _, err = dal.AddTagArticle(&tagArticle); err != nil {
		return
	}

	return
}

// 获取图片
func GetImageData(imageUrl string, articleId string) (err error) {
	has, err := dal.HasImage(imageUrl, articleId)
	if err != nil {
		return err
	}

	if has {
		return nil
	}

	httpReq, err := imgRequest(&httpRequest.HttpRequest{
		Url:    imageUrl,
		Method: "GET",
	})
	if err != nil {
		return err
	}

	result, err := httpReq.DoRequest()
	if err != nil {
		return err
	}

	_, err = dal.AddImage(&model.Image{
		ArticleId: articleId,
		Url:       imageUrl,
		Code:      result.Resp.StatusCode,
		Ctype:     result.Resp.Header.Get("content-type"),
		Data:      *result.Data,
	})
	if err != nil {
		return err
	}

	return nil

}
