package file

import "github.com/spf13/cobra"

const (
	DaFile = "file"
	DaPath = "filestore.path"
)

func AddFlags(cmd *cobra.Command) {
	cmd.Flags().String(DaPath, DefaultFileStoreCfg().Directory, "filestore path")
}

func ParseConfig(cmd *cobra.Command) Config {
	cfg := DefaultFileStoreCfg()
	if path := cmd.Flag(DaPath).Value.String(); path != "" {
		cfg.Directory = path
	}

	return cfg
}
