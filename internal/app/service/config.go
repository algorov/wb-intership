package service

import (
	"l0Service/internal/app/nats"
	"l0Service/internal/app/store"
)

// Config ...
type Config struct {
	BindAddr      string        `toml:"bind_addr"`
	LogLevel      string        `toml:"log_level"`
	Store         *store.Config `toml:"store"`
	NatsStreaming *nats.Config  `toml:"nats_streaming"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:      ":8080",
		LogLevel:      "debug",
		Store:         store.NewConfig(),
		NatsStreaming: nats.NewConfig(),
	}
}
