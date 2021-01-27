package main

import (
	"fmt"
	"juejinCollections/collectReq"
	"juejinCollections/config"
	"juejinCollections/logger"
	"juejinCollections/middleware"
	"juejinCollections/model"

	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	dal "juejinCollections/dal"
)

func main() {
	defer func() {
		logger.Logger.Error("???")
	}()
	// r := gin.Default()
	conf := config.Config

	if conf.Debug {
		logger.SetDebugLog(true)
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	dal.NewDal()

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery(), middleware.Recovery())

	// r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
	// 	if err, ok := recovered.(string); ok {
	// 		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	// 	}
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// }))

	noRoute := func(c *gin.Context) {
		if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status": false,
				"msg":    "NOT FOUND",
			})
		}
	}

	r.LoadHTMLGlob("frontend/dist/*.html")
	r.NoRoute(noRoute)
	r.NoMethod(noRoute)

	r.Static("/_assets", "frontend/dist/_assets")
	r.StaticFile("/favicon.ico", "frontend/dist/favicon.ico")

	imageGroup := r.Group("/images")
	imageGroup.GET("/article/:articleId", func(c *gin.Context) {
		fmt.Println("??")
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
	})

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

	r.Any("/", noRoute)

	logger.ExitHook.Add(func() {
		if dal.DbDal.Engine != nil {
			dal.DbDal.Engine.Close()
		}
	})

	go collectReq.Run()
	// if !conf.Debug {
	// 	go collectReq.Run()
	// }
	r.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
