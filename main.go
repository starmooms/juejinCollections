package main

import (
	"fmt"
	"juejinCollections/collectReq"
	"juejinCollections/middleware"

	// "juejinCollections/httpRequest"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "juejinCollections/dal"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recovery())

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
		panic("?")
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": true,
		// 	"data": gin.H{
		// 		"a": 1,
		// 		"b": 2,
		// 	},
		// })
	})

	// https://api.juejin.cn/interact_api/v1/collectionSet/list
	// 1116759544852221
	// 2664871913078168
	// httpReq := httpRequest.Request(&httpRequest.HttpRequest{
	// 	Url:    "http://www-test.yingsheng.com/webhome/api/operate",
	// 	Method: "GET",
	// 	Params: &gin.H{
	// 		"user_id": 1116759544852221,
	// 		"cursor":  0,
	// 		"limit":   20,
	// 	},
	// })
	// data, err := httpReq.DoRequest()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)

	// httpReq2 := httpRequest.Request(&httpRequest.HttpRequest{
	// 	Url:    "http://www-test.yingsheng.com/webhome/api/operate",
	// 	Method: "POST",
	// })
	// data2, err2 := httpReq2.DoRequest()
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// fmt.Println(data2)

	collectReq.GetTagList()

	// httpRequest.Get("http://www-test.yingsheng.com/webhome/api/operate", &gin.H{
	// 	"user_id": 1116759544852221,
	// 	"cursor":  0,
	// 	"limit":   20,
	// })

	// r.GET("/*url", func(c *gin.Context) {
	// 	// c.JSON(200, gin.H{
	// 	// 	"message": "pong",
	// 	// })
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })
	r.Run("localhost:8012") // listen and serve on 0.0.0.0:8080
}
