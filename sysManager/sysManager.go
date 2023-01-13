package sysManager

import (
	"juejinCollections/config"
	"juejinCollections/logger"
	"juejinCollections/statikFs"

	"github.com/getlantern/systray"
	"juejinCollections/server"
)

var srv *server.Server

func createServer() *server.Server {
	conf := config.Config
	srv = &server.Server{
		Port:   conf.Port,
		Host:   conf.Host,
		Logger: logger.Logger,
	}
	go srv.Start()
	return srv
}

func onReady() {
	iconData := statikFs.GetFileDataMust("./frontend/dist/favicon.ico")

	systray.SetIcon(iconData)
	systray.SetTitle("juejin-collections-app")
	systray.SetTooltip("juejin-collections")

	mOpenBrowser := systray.AddMenuItem("打开", "打开页面")
	mQuit := systray.AddMenuItem("退出", "关闭当前进程")

	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
		case <-mOpenBrowser.ClickedCh:
			if srv != nil {
				srv.OpenBrowser()
			}
		}
	}
}

func onExit() {
	if srv != nil {
		srv.Exit()
		srv = nil
	}
}

func Init() {
	createServer()
	systray.Run(onReady, onExit)
}
