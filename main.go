package main

import (
	"fmt"
	"juejinCollections/collectReq"
	"juejinCollections/config"
	"juejinCollections/logger"
	"juejinCollections/middleware"

	// "juejinCollections/httpRequest"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"

	dal "juejinCollections/dal"
)

func main() {
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

	r.LoadHTMLGlob("frontend/dist/*.html")
	r.NoRoute(func(c *gin.Context) {
		fmt.Println("end")
		if c.Request.Method == "GET" {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status": false,
				"msg":    "NOT FOUND",
			})
		}

	})

	r.Static("/_assets", "frontend/dist/_assets")
	r.StaticFile("/favicon.ico", "frontend/dist/favicon.ico")

	r.POST("/api/abc", func(c *gin.Context) {
		// panic(aE.New("??"))
		panic(errors.New("??"))

		// c.JSON(http.StatusOK, gin.H{
		// 	"status": true,
		// 	"data": gin.H{
		// 		"a": 1,
		// 		"b": 2,
		// 	},
		// })
	})

	
	collectReq.Run()
	r.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
