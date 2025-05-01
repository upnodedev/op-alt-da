package config

import (
	"os"

	"github.com/spf13/viper"
)

// env constants
const (
	PlasmaDaHttpHost = "HTTP_HOST"
	PlasmaDaHttpPort = "HTTP_PORT"
	PlasmaDaType     = "DA_TYPE"
	PlasmaDaId       = "DA_ID"
	PlasmaDaHomeDir  = "HOME_DIR"
	PlasmaEvmRpcUrl  = "EVM_RPC_URL"
	PlasmaPrivateKey = "PRIVATE_KEY"
	PlasmaChainId    = "CHAIN_ID"
	PlasmaHubAddr    = "ALT_DA_HUB_ADDR"
)

type App struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DA           string `json:"da"`
	DaID         string `json:"da_id"`
	HomeDir      string `json:"home_dir"`
	EvmRpcUrl    string `json:"evm_rpc_url"`
	PrivateKey   string `json:"private_key"`
	ChainId      int64  `json:"chain_id"`
	AltDaHubAddr string `json:"alt_da_hub_addr"`
}

func DefaultConfig() App {
	// default app config will read from the environment variables
	// if set by flag, it will override the default values
	homeDir, _ := os.UserHomeDir()
	cfg := App{
		Host:         "localhost",
		Port:         8087,
		DA:           "file",
		DaID:         "0x000c", // it is celestia
		HomeDir:      homeDir,
		EvmRpcUrl:    "https://mainnet.optimism.io",
		PrivateKey:   "",
		ChainId:      10,
		AltDaHubAddr: "0x2F77fDf77E5a13092D08028188B40b691c41FbDe",
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
	if daId := viper.GetString(PlasmaDaId); daId != "" {
		cfg.DaID = daId
	}
	if evmRpcUrl := viper.GetString(PlasmaEvmRpcUrl); evmRpcUrl != "" {
		cfg.EvmRpcUrl = evmRpcUrl
	}
	if privateKey := viper.GetString(PlasmaPrivateKey); privateKey != "" {
		cfg.PrivateKey = privateKey
	}
	if chainId := viper.GetInt64(PlasmaChainId); chainId > 0 {
		cfg.ChainId = chainId
	}
	if plasmaHubAddr := viper.GetString(PlasmaHubAddr); plasmaHubAddr != "" {
		cfg.AltDaHubAddr = plasmaHubAddr
	}

	return cfg
}

//func init() {
//	viper.AutomaticEnv()
//	viper.SetEnvPrefix(PrefixEnv)
//}
