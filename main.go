package main

import (
	"juejinCollections/collectReq"
	"juejinCollections/config"
	"juejinCollections/logger"
	"juejinCollections/statikFs"
	"juejinCollections/sysManager"

	"github.com/gin-gonic/gin"

	dal "juejinCollections/dal"
)

func main() {
	defer func() {
		logger.Logger.Debug("All exit")
	}()
	// r := gin.Default()
	conf := config.Config
	statikFs.InitStatikFs()

	if conf.IsDebug {
		logger.SetDebugLog(true)
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	dal.NewDal(conf.DbFile)
	collectReq.InitCollectReq()

	sysManager.Init()
}
