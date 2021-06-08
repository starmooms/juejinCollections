package server

import (
	"juejinCollections/server/controller"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	r.LoadHTMLGlob("frontend/dist/*.html")
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "GET" {
			s := strings.ToUpper(c.Request.Header.Get("X-Requested-With"))
			isXhr := s == "XMLHTTPREQUEST"
			if !isXhr {
				c.HTML(http.StatusOK, "index.html", gin.H{})
				return
			}
		}

		code := http.StatusNotFound
		c.JSON(code, gin.H{
			"status": false,
			"msg":    http.StatusText(code),
		})
	})

	r.Static("/assets", "frontend/dist/assets")
	r.StaticFile("/favicon.ico", "frontend/dist/favicon.ico")

	imageGroup := r.Group("/images")
	{
		imageGroup.GET("/article/:articleId", controller.ArticleImage)
	}

	r.POST("/api/abc", func(c *gin.Context) {
		// panic(aE.New("??"))
		// panic(errors.New("??"))
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": gin.H{
				"a": 1,
				"b": 2,
			},
		})
	})

	api := r.Group("/api")
	{
		api.GET("/getArticle", controller.GetArticle)
	}

}
