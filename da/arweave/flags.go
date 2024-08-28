package arweave

import "github.com/spf13/cobra"

const (
	arClientUrl  = "ar.client_url"
	arWalletPath = "ar.wallet_path"
)

func AddFlags(cmd *cobra.Command) {
	cmd.Flags().String(arClientUrl, DefaultArConfig().ClientUrl, "Arweave client url")
	cmd.Flags().String(arWalletPath, DefaultArConfig().WalletPath, "Arweave wallet path")
}

func ParseConfig(cmd *cobra.Command) Config {
	cfg := DefaultArConfig()
	if clientUrl := cmd.Flag(arClientUrl).Value.String(); clientUrl != "" {
		cfg.ClientUrl = clientUrl
	}

	if walletPath := cmd.Flag(arWalletPath).Value.String(); walletPath != "" {
		cfg.WalletPath = walletPath
	}

	return cfg
}
