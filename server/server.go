package server

import (
	"fmt"
	"github.com/Master0fMagic/rest-health-check-service/config"
	"github.com/Master0fMagic/rest-health-check-service/health"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Server struct {
	srv *gin.Engine
	cfg config.HTTPConfig
}

func New(cfg config.HTTPConfig, h health.URLHealthCheck, logger log.FieldLogger) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(ginlogrus.Logger(logger))

	r.POST("/urls-health", getUrlsHealthRequestHandlerFunc(h))

	return &Server{
		srv: r,
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	log.Infof("running http server on port %d", s.cfg.Port)
	return s.srv.Run(fmt.Sprintf(":%d", s.cfg.Port))
}
