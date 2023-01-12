package server

import (
	"io/ioutil"
	"juejinCollections/server/controller"
	"juejinCollections/server/statikFs"
	"juejinCollections/tool"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	statikFs.InitStatikFs()

	// r.LoadHTMLGlob("frontend/dist/*.html")
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "GET" {
			s := strings.ToUpper(c.Request.Header.Get("X-Requested-With"))
			isXhr := s == "XMLHTTPREQUEST"
			if !isXhr {
				// 模板的方式参考 https://github.com/rakyll/statik/issues/18
				f, err := statikFs.GetFileSystem().Open("/index.html")
				if err != nil {
					tool.PanicErrMsg("no index.html")
				}
				b, err := ioutil.ReadAll(f)
				c.Data(http.StatusOK, "text/html;charset=utf-8", b)
				// c.HTML(http.StatusOK, "index.html", gin.H{})
				return
			}
		}

		code := http.StatusNotFound
		c.JSON(code, gin.H{
			"status": false,
			"msg":    http.StatusText(code),
		})
	})

	// r.Static("/assets", "frontend/dist/assets")
	// r.StaticFile("/favicon.ico", "frontend/dist/favicon.ico")

	// statikFs.GetFile()

	r.StaticFS("/assets", statikFs.OpenDir("/assets/"))
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("/favicon.ico", statikFs.GetFileSystem())
	})

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
	}

}
