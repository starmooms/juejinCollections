package app

import (
	"errors"
	"juejinCollections/tool"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseData struct {
	data   interface{}
	code   int
	msg    string
	status bool
	err    error
}

func (g *Gin) Response(r *ResponseData) {
	if r.err != nil {
		tool.ShowErr(r.err)
	}

	g.C.JSON(r.code, &gin.H{
		"status": r.status,
		"data":   r.data,
		"msg":    r.msg,
	})
	return
}

func (g *Gin) BackData(data interface{}) {
	result := &ResponseData{
		code:   http.StatusOK,
		status: true,
		data:   data,
	}
	g.Response(result)
	return
}

func (g *Gin) BackParamsErr(err error) {
	if err == nil {
		g.BackErr(errors.New("BackParamsErr err is must no null"))
		return
	}

	errs := err.(validator.ValidationErrors)[0]
	result := &ResponseData{
		code:   http.StatusBadRequest,
		msg:    errs.Translate(Trans),
		status: false,
	}
	g.Response(result)
	return
}

func (g *Gin) BackMessageErr(msg string) {
	result := &ResponseData{
		code:   http.StatusBadRequest,
		msg:    msg,
		status: false,
	}
	g.Response(result)
	return
}

func (g *Gin) BackErr(err error) {
	result := &ResponseData{
		code:   http.StatusInternalServerError,
		err:    err,
		msg:    err.Error(),
		status: false,
	}
	g.Response(result)
	return
}

func (g *Gin) BackErrCode(code int) {
	result := &ResponseData{
		code:   code,
		msg:    http.StatusText(code),
		status: false,
	}
	g.Response(result)
	return
}
