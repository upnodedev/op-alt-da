package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

var DefaultConfigPath = "config"

const (
	HttpHost = "server.http_host"
	HttpPort = "server.http_port"
	DA       = "server.da"
)

type App struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	DA   string `json:"da"`
}

func DefaultConfig() App {
	return App{
		Host: "localhost",
		Port: 8087,
		DA:   "file",
	}
}

func NewAppConfig(homeDir string) App {
	loadConfig(homeDir)
	cfg := DefaultConfig()

	if host := viper.GetString(HttpHost); host != "" {
		cfg.Host = host
	}
	if port := viper.GetInt(HttpPort); port > 0 {
		cfg.Port = port
	}
	if da := viper.GetString(DA); da != "" {
		cfg.DA = da
	}

	return cfg
}

func loadConfig(homeDir string) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath(filepath.Join(homeDir, DefaultConfigPath))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %s", err))
	}

	viper.WatchConfig()
	viper.SetEnvPrefix("CS")
}
