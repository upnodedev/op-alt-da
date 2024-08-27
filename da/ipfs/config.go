package ipfs

import "github.com/spf13/viper"

const (
	PlasmaIpfsUrl = "IPFS_URL"
)

type Config struct {
	Url string
}

func DefaultIpfsConfig() Config {
	cfg := Config{
		Url: "localhost:5001",
	}

	if url := viper.GetString(PlasmaIpfsUrl); url != "" {
		cfg.Url = url
	}

	return cfg
}
