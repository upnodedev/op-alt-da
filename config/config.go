package config

import (
	"github.com/spf13/viper"
)

const (
	HttpPort = "HTTP_PORT"
	HttpHost = "HTTP_HOST"
)

type App struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func DefaultConfig() App {
	return App{
		Host: "localhost",
		Port: 8087,
	}
}

func NewAppConfig() App {
	cfg := DefaultConfig()

	if host := viper.GetString(HttpHost); host != "" {
		cfg.Host = host
	}
	if port := viper.GetInt(HttpPort); port > 0 {
		cfg.Port = port
	}

	return cfg
}
