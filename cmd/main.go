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
	cfgApp := config.NewAppConfig()
	store, err := da.NewCelestiaStore(da.CelestiaCfg{
		Rpc:                 "localhost",
		AuthToken:           "",
		Namespace:           "",
		EthFallbackDisabled: false,
		MaxBlobSize:         0,
		GasPrice:            0,
	})
	if err != nil {
		panic(err)
	}
	daServer := plasma.NewDAServer(cfgApp, store, slog.Default())
	daServer.Start()
}
