package server

import (
	"context"
	"fmt"
	"juejinCollections/config"
	"juejinCollections/server/middleware"
	"juejinCollections/server/websocket"
	"juejinCollections/tool"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Logger   *logrus.Logger
	Host     string
	Port     int
	Engine   *gin.Engine
	srv      *http.Server
	hostName string
}

func (s *Server) getServerUrl() string {
	return fmt.Sprintf("http://%s", s.hostName)
}

func (s *Server) afterStart() {
	serverUrl := s.getServerUrl()
	fmt.Printf("server start %s", serverUrl)
	if !config.Config.IsDevelopment {
		s.OpenBrowser()
	}
}

func (s *Server) OpenBrowser() {
	tool.OpenBrowser(s.getServerUrl())
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

	s.hostName = fmt.Sprintf("%s:%d", s.Host, s.Port)

	// serverUrl := fmt.Sprintf("http://%s", hostName)
	// go func() {
	// 	fmt.Printf("server start %s", serverUrl)
	// 	if !config.Config.IsDevelopment {
	// 		tool.OpenBrowser(serverUrl)
	// 	}
	// }()
	// r.Run(hostName)

	s.srv = &http.Server{
		Addr:    s.hostName,
		Handler: r,
	}

	go s.afterStart()
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		tool.ShowErr(err)
	}
}

func (s *Server) Exit() {
	s.Logger.Debug("start server exiting...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		tool.ShowErr(err)
	}

	select {
	case <-ctx.Done():
		s.Logger.Debug("server Exit done")
	}
	s.Logger.Debug("server exiting")
	s.srv = nil
}
