package tool

import (
	"juejinCollections/logger"
	"reflect"

	"github.com/cockroachdb/errors"
)

var logs = logger.Logger

func PanicErr(err error) {
	if err != nil {
		logs.Errorf("%+v", errors.NewWithDepth(1, err.Error()))
		panic(err)
	}
}

/** 判断是否有错误堆载，没有添加错误堆载后返回 */
func SetErrStack(err error, depth ...int) error {
	rt := reflect.TypeOf(err)
	kind := rt.Kind()
	hasStack := false
	if kind == reflect.Ptr {
		rt = rt.Elem()
		kind = rt.Kind()
	}
	if kind == reflect.Struct {
		_, hasStack = rt.FieldByName("stack")
	}
	if !hasStack {
		setDepth := 1
		if len(depth) >= 1 {
			setDepth = depth[0]
		}
		return errors.NewWithDepth(setDepth, err.Error())
	}
	return err
}

/** 返回并打印错误 */
func BackError(err error) error {
	err = SetErrStack(err, 2)
	logs.Errorf("%+v\n", err)
	return err
}

func BackNewError(msg string, depth ...int) error {
	setDepth := 1
	if len(depth) >= 1 {
		setDepth = depth[0]
	}
	err := errors.NewWithDepth(setDepth, msg)
	logs.Errorf("%+v\n", err)
	return err
}

/** 返回并打印错误 */
func ShowErr(err error) error {
	err = SetErrStack(err, 2)
	logs.Errorf("%+v\n", err)
	return err
}
