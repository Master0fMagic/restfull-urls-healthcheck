package config

import "time"

type Config struct {
	LogLevel          string            `env:"LOG_LEVEL" envDefault:"DEBUG"`
	HttpConfig        HttpConfig        `envPrefix:"HTTP_"`
	HealthCheckConfig HealthCheckConfig `envPrefix:"HEALTHCHECK_"`
}

type HttpConfig struct {
	Port int `env:"PORT" envDefault:"8080"`
}

type HealthCheckConfig struct {
	Timeout                 time.Duration `env:"TIMEOUT" envDefault:"2s"`
	StopOnFailure           bool          `env:"STOP_ON_FAILURE" envDefault:"false"`        // if true, will cancel ctx on first inactive API
	MaxProcessingGoroutines int           `env:"MAX_PROCESSING_GOROUTINES" envDefault:"-1"` // set to -1 to disable limit
}
