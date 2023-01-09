package controller

import (
	"juejinCollections/collectReq"
	"juejinCollections/dal"
	"juejinCollections/model"
	"juejinCollections/server/app"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// 获取图片
func ArticleImage(c *gin.Context) {
	url, err := url.PathUnescape(c.Query("url"))
	if err != nil {
		c.Error(err)
		return
	}

	image := &model.Image{
		Url:       url,
		ArticleId: c.Param("articleId"),
	}
	dal.GetImage(image)
	if image.Code == 0 {
		image.Code = 404
	}
	c.Data(image.Code, image.Ctype, image.Data)
}

// 获取文章
func GetArticle(c *gin.Context) {
	appG := &app.Gin{
		C: c,
	}
	var err error
	params := &GetArticleParams{}
	err = c.ShouldBindQuery(params)
	if err != nil {
		appG.BackParamsErr(err)
		return
	}

	article := &model.Article{
		ArticleId: params.ArticleId,
	}
	has, err := dal.Get(article)

	if appG.HasError(err) {
		return
	}

	if !has {
		appG.BackErrCode(http.StatusNotFound)
		return
	}

	appG.BackData(article)
}

// 同步收藏集
func RunSyncCollection(c *gin.Context) {
	appG := &app.Gin{
		C: c,
	}

	if collectReq.HasRunAction {
		appG.BackMessageErr("当前正在同步收藏集，不能重复触发")
		return
	}

	go collectReq.Run()
	appG.BackData(nil)
}
