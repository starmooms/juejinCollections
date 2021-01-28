package server

import (
	"fmt"
	"juejinCollections/server/middleware"

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

	r.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
