package arweave

import "os"

const (
	PlasmaArClientUrl  = "AR_CLIENT_URL"
	PlasmaArWalletPath = "AR_WALLET_PATH"
)

type Config struct {
	ClientUrl  string
	WalletPath string
}

func DefaultArConfig() Config {
	cfg := Config{
		ClientUrl:  "https://arweave.net",
		WalletPath: "",
	}

	if clientUrl := os.Getenv(PlasmaArClientUrl); clientUrl != "" {
		cfg.ClientUrl = clientUrl
	}

	if walletPath := os.Getenv(PlasmaArWalletPath); walletPath != "" {
		cfg.WalletPath = walletPath
	}

	return cfg
}
