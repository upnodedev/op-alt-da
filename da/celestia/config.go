package celestia

import "github.com/spf13/viper"

const (
	PlasmaCelestiaRpc                 = "CELESTIA_RPC"
	PlasmaCelestiaAuthToken           = "CELESTIA_AUTH_TOKEN"
	PlasmaCelestiaNamespace           = "CELESTIA_NAMESPACE"
	PlasmaCelestiaEthFallbackDisabled = "CELESTIA_ETH_FALLBACK_DISABLED"
	PlasmaCelestiaMaxBlobSize         = "CELESTIA_MAX_BLOB_SIZE"
	PlasmaCelestiaGasPrice            = "CELESTIA_GAS_PRICE"
)

type Config struct {
	Rpc                 string
	AuthToken           string
	Namespace           string
	EthFallbackDisabled bool
	MaxBlobSize         uint64
	GasPrice            float64
}

func DefaultCelestiaConfig() Config {
	cfg := Config{
		Rpc:                 "http://localhost:7980",
		AuthToken:           "",
		Namespace:           "",
		EthFallbackDisabled: false,
		MaxBlobSize:         2000,
		GasPrice:            0.002,
	}

	if rpc := viper.GetString(PlasmaCelestiaRpc); rpc != "" {
		cfg.Rpc = rpc
	}
	if authToken := viper.GetString(PlasmaCelestiaAuthToken); authToken != "" {
		cfg.AuthToken = authToken
	}
	if namespace := viper.GetString(PlasmaCelestiaNamespace); namespace != "" {
		cfg.Namespace = namespace
	}
	if ethFallbackDisabled := viper.GetBool(PlasmaCelestiaEthFallbackDisabled); ethFallbackDisabled {
		cfg.EthFallbackDisabled = ethFallbackDisabled
	}
	if maxBlobSize := viper.GetUint64(PlasmaCelestiaMaxBlobSize); maxBlobSize > 0 {
		cfg.MaxBlobSize = maxBlobSize
	}
	if gasPrice := viper.GetFloat64(PlasmaCelestiaGasPrice); gasPrice > 0 {
		cfg.GasPrice = gasPrice
	}

	return cfg
}
