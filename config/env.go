package config

import "sync"

type Config struct {
	Port    string
	BaseURL string
	Mode    string
}

var (
	envConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		envConfig = &Config{
			Port:    "5000",
			BaseURL: "http://localhost:5000",
			Mode:    "development",
		}
	})

	return envConfig
}
