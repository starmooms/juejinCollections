package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator

type Gin struct {
	C *gin.Context
}

func init() {
	Trans = ValidatorTranslator()
}

// // gin 设置跨中间件传递
// func (g *Gin) CheckLogin(data interface{}) {
// 	g.C.Set("userInfo", &gin.H{
// 		"name": "...",
// 	})
// 	return
// }

func (g *Gin) HasError(err error) bool {
	if err != nil {
		g.BackErr(err)
		return true
	}
	return false
}
