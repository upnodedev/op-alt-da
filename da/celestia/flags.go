package celestia

import "github.com/spf13/cobra"

// Flags for Celestia.
const (
	Rpc                 = "celestia.rpc"
	AuthToken           = "celestia.auth_token"
	Namespace           = "celestia.namespace"
	EthFallbackDisabled = "celestia.eth_fallback_disabled"
	MaxBlobSize         = "celestia.max_blob_size"
	GasPrice            = "celestia.gas_price"
)

func AddFlags(cmd *cobra.Command) {
	cmd.Flags().String(Rpc, DefaultCelestiaConfig().Rpc, "Celestia RPC port")
	cmd.Flags().String(AuthToken, DefaultCelestiaConfig().AuthToken, "Celestia auth token")
	cmd.Flags().String(Namespace, DefaultCelestiaConfig().Namespace, "Celestia namespace")
	cmd.Flags().Bool(EthFallbackDisabled, DefaultCelestiaConfig().EthFallbackDisabled, "Disable Ethereum fallback")
	cmd.Flags().Uint64(MaxBlobSize, DefaultCelestiaConfig().MaxBlobSize, "Max blob size")
	cmd.Flags().Float64(GasPrice, DefaultCelestiaConfig().GasPrice, "Gas price")
}

func ParseConfig(cmd *cobra.Command) Config {
	cfg := DefaultCelestiaConfig()
	if rpc := cmd.Flag(Rpc).Value.String(); rpc != "" {
		cfg.Rpc = rpc
	}
	if authToken := cmd.Flag(AuthToken).Value.String(); authToken != "" {
		cfg.AuthToken = authToken
	}
	if namespace := cmd.Flag(Namespace).Value.String(); namespace != "" {
		cfg.Namespace = namespace
	}
	if ethFallbackDisabled, err := cmd.Flags().GetBool(EthFallbackDisabled); err == nil {
		cfg.EthFallbackDisabled = ethFallbackDisabled
	}
	if maxBlobSize, err := cmd.Flags().GetUint64(MaxBlobSize); err == nil {
		cfg.MaxBlobSize = maxBlobSize
	}

	return cfg
}
