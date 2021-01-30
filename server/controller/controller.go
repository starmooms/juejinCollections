package controller

import (
	"juejinCollections/dal"
	"juejinCollections/model"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

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

func GetArticle(c *gin.Context) {
	var err error
	params := &GetArticleParams{}
	err = c.ShouldBindQuery(params)
	if err != nil {
		BackParamsErr(c, err)
		return
	}

	article := &model.Article{
		ArticleId: params.ArticleId,
	}
	has, err := dal.Get(article)
	if !has {
		code := http.StatusNotFound
		c.JSON(code, &gin.H{
			"status": false,
			"msg":    http.StatusText(code),
		})
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"status": true,
		"data": &gin.H{
			"article": article,
		},
	})
}
