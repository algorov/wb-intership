package service

import "l0Service/internal/app/store"

// Config ...
type Config struct {
	BindAddr      string `toml:"bind_addr"`
	LogLevel      string `toml:"log_level"`
	NatsClusterId string `toml:"nats_cluster_id"`
	NatsClientId  string `toml:"nats_client_id"`
	NatsTopic     string `toml:"nats_topic"`
	DurableName   string `toml:"durable_name"`
	Store         *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:      ":8080",
		LogLevel:      "debug",
		NatsClusterId: "test-cluster",
		NatsClientId:  "slave",
		NatsTopic:     "main",
		DurableName:   "defer",
		Store:         store.NewConfig(),
	}
}
