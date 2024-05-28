package config

import "github.com/spf13/viper"

const HttpPort = "HTTP_PORT"

type App struct {
	HttpPort string `json:"http_port"`
}

func DefaultConfig() App {
	return App{
		HttpPort: "8087",
	}
}

func NewAppConfig() App {
	cfg := DefaultConfig()

	if port := viper.GetString(HttpPort); port != "" {
		cfg.HttpPort = port
	}

	return cfg
}
