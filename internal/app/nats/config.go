package nats

// Config ...
type Config struct {
	NatsClusterId string `toml:"nats_cluster_id"`
	NatsClientId  string `toml:"nats_client_id"`
	NatsTopic     string `toml:"nats_topic"`
	DurableName   string `toml:"durable_name"`
}

func NewConfig() *Config {
	return &Config{
		NatsClusterId: "test-cluster",
		NatsClientId:  "slave",
		NatsTopic:     "main",
		DurableName:   "defer",
	}
}
