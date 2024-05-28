package main

import (
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"plasma"
	"plasma/config"
	"plasma/da"
)

var Version = "v0.0.1"

func main() {
	viper.AutomaticEnv()
	l := slog.Default()
	cfgApp := config.NewAppConfig()
	store := da.NewFileStore("data.json")
	daServer := plasma.NewDAServer(cfgApp, store, l)
	daServer.Start()
}
