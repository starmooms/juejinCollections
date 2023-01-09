package app

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// https://studygolang.com/articles/28414?fr=sidebar
func ValidatorTranslator() ut.Translator {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		// //注册自定义函数
		// _ = v.RegisterValidation("bookabledate", bookableDate)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			label := fld.Tag.Get("label")
			if label != "" {
				return label
			}

			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// //根据提供的标记注册翻译
		// v.RegisterTranslation("bookabledate", trans, func(ut ut.Translator) error {
		// 	return ut.Add("bookabledate", "{0}不能早于当前时间或{1}格式错误!", true)
		// }, func(ut ut.Translator, fe validator.FieldError) string {
		// 	t, _ := ut.T("bookabledate", fe.Field(), fe.Field())
		// 	return t
		// })
	}
	return trans
}
