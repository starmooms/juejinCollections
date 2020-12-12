package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("frontend/dist/*.html")

	p, _ := filepath.Abs("./1.txt")
	fmt.Println(p)

	// distPath, _ := filepath.Abs("./frontend/dist")
	// fmt.Println(distPath)

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
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": gin.H{
				"a": 1,
				"b": 2,
			},
		})
	})

	// r.GET("/*url", func(c *gin.Context) {
	// 	// c.JSON(200, gin.H{
	// 	// 	"message": "pong",
	// 	// })
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })
	r.Run("localhost:8012") // listen and serve on 0.0.0.0:8080
}
