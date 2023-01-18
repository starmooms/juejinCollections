package server

import (
	"juejinCollections/server/controller"
	"juejinCollections/statikFs"
	"juejinCollections/tool"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func isXhr(c *gin.Context) bool {
	s := strings.ToUpper(c.Request.Header.Get("X-Requested-With"))
	return s == "XMLHTTPREQUEST"
}

func SetRoute(r *gin.Engine) {

	// 如果需要模板的方式参考 https://github.com/rakyll/statik/issues/18
	// r.LoadHTMLGlob("frontend/dist/*.html")

	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "GET" && !isXhr(c) {
			data, err := statikFs.GetFileData("./frontend/dist/index.html")
			if err != nil {
				tool.PanicErr(err)
			}
			c.Data(http.StatusOK, "text/html;charset=utf-8", data)
			// c.HTML(http.StatusOK, "index.html", gin.H{})
			return
		}

		code := http.StatusNotFound
		c.JSON(code, gin.H{
			"status": false,
			"msg":    http.StatusText(code),
		})
	})

	statikFs.SetGinStatic(r, "/assets", "frontend/dist/assets")
	statikFs.SetGinStaticFile(r, "/favicon.ico", "frontend/dist/favicon.ico")

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
		api.POST("/syncCollection", controller.RunSyncCollection)
		api.GET("/searchArticle", controller.SearchArticle)
	}

}
