package apiserver

import "os"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	port, exists := os.LookupEnv("$PORT")
	if !exists {
		port = "8080"
	}
	return &Config{
		BindAddr: ":" + port,
		LogLevel: "debug",
	}
}
