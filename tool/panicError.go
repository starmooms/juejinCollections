package tool

import (
	"juejinCollections/logger"

	"github.com/cockroachdb/errors"
)

func PanicErr(err error) {
	if err != nil {
		logs := logger.GetLog()
		logs.Errorf("%+v", errors.NewWithDepth(1, err.Error()))
		panic(err)
	}
}
