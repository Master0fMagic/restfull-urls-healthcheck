package main

import (
	"github.com/Master0fMagic/rest-health-check-service/config"
	"github.com/Master0fMagic/rest-health-check-service/health"
	"github.com/Master0fMagic/rest-health-check-service/server"
	"github.com/caarlos0/env/v10"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		log.WithError(err).Fatal("error reading config")
	}

	initLogger(cfg.LogLevel)

	srv := server.New(cfg.HttpConfig,
		health.New(cfg.HealthCheckConfig),
		log.New(),
	)

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("error running server")
	}
}

func initLogger(logLevel string) {
	lvl, err := log.ParseLevel(logLevel)
	if err != nil {
		lvl = log.InfoLevel
		log.WithError(err).Warn("error parsing log level. Log level set to INFO")
	}
	log.SetLevel(lvl)
}
