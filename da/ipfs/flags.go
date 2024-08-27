package ipfs

import "github.com/spf13/cobra"

const ipfsUrl = "ipfs.url"

func AddFlags(cmd *cobra.Command) {
	cmd.Flags().String(ipfsUrl, DefaultIpfsConfig().Url, "IPFS https url")
}

func ParseConfig(cmd *cobra.Command) Config {
	cfg := DefaultIpfsConfig()
	if url := cmd.Flag(ipfsUrl).Value.String(); url != "" {
		cfg.Url = url
	}

	return cfg
}
