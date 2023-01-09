package main

import (
	"juejinCollections/config"
	"juejinCollections/logger"
	"juejinCollections/server"

	"github.com/gin-gonic/gin"

	dal "juejinCollections/dal"
)

func main() {
	defer func() {
		logger.Logger.Error("???")
	}()
	// r := gin.Default()
	conf := config.Config

	if conf.IsDebug {
		logger.SetDebugLog(true)
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	dal.NewDal(conf.DbFile)

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
