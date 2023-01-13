package main

import (
	"juejinCollections/collectReq"
	"juejinCollections/config"
	"juejinCollections/logger"
	"juejinCollections/server"
	"juejinCollections/statikFs"

	"github.com/gin-gonic/gin"

	dal "juejinCollections/dal"
)

func main() {
	defer func() {
		logger.Logger.Error("exit")
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

	// go collectReq.Run()
	// if !conf.Debug {
	// 	go collectReq.Run()
	// }
	srv := &server.Server{
		Port:   conf.Port,
		Host:   conf.Host,
		Logger: logger.Logger,
	}
	srv.Start()

}
