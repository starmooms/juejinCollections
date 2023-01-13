package server

import (
	"fmt"
	"juejinCollections/config"
	"juejinCollections/server/middleware"
	"juejinCollections/server/websocket"
	"juejinCollections/tool"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Logger *logrus.Logger
	Host   string
	Port   int
	Engine *gin.Engine
}

func (s *Server) Start() {
	r := gin.New()
	s.Engine = r

	r.Use(
		middleware.Logger(s.Logger),
		gin.Recovery(),
		middleware.Recovery(s.Logger),
	)
	SetRoute(r)
	websocket.Start(r)

	hostName := fmt.Sprintf("%s:%d", s.Host, s.Port)
	serverUrl := fmt.Sprintf("http://%s", hostName)

	go func() {
		fmt.Printf("server start %s", serverUrl)
		if !config.Config.IsDevelopment {
			tool.OpenBrowser(serverUrl)
		}
	}()

	r.Run(hostName)
}
