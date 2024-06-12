package config

import (
	"github.com/spf13/viper"
	"os"
)

// env constants
const (
	PlasmaDaHttpHost = "HTTP_HOST"
	PlasmaDaHttpPort = "HTTP_PORT"
	PlasmaDaType     = "DA_TYPE"
	PlasmaDaHomeDir  = "HOME_DIR"
)

type App struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	DA      string `json:"da"`
	HomeDir string `json:"home_dir"`
}

func DefaultConfig() App {
	// default app config will read from the environment variables
	// if set by flag, it will override the default values
	homeDir, _ := os.UserHomeDir()
	cfg := App{
		Host:    "localhost",
		Port:    8087,
		DA:      "file",
		HomeDir: homeDir,
	}

	if homeDir := viper.GetString(PlasmaDaHomeDir); homeDir != "" {
		cfg.HomeDir = homeDir
	}
	if host := viper.GetString(PlasmaDaHttpHost); host != "" {
		cfg.Host = host
	}
	if port := viper.GetInt(PlasmaDaHttpPort); port > 0 {
		cfg.Port = port
	}
	if da := viper.GetString(PlasmaDaType); da != "" {
		cfg.DA = da
	}

	return cfg
}

//func init() {
//	viper.AutomaticEnv()
//	viper.SetEnvPrefix(PrefixEnv)
//}
