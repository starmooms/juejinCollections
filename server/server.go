package server

import (
	"fmt"
	"juejinCollections/server/middleware"
	"juejinCollections/server/websocket"

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

	fmt.Printf("server start http://%s:%d", s.Host, s.Port)
	r.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))

	// r.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
	// websocket.Start(r, s.Port)
}
