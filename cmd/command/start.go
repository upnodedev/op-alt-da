package command

import (
	"errors"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"plasma"
	"plasma/config"
	"plasma/da"
)

func StartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir := cmd.Flag("home").Value.String()
			cfgApp := config.NewAppConfig(homeDir)

			var store da.KVStore
			var err error
			switch cfgApp.DA {
			case da.DaFile:
				cfgFileStore := da.DefaultFileStoreCfg()
				store = da.NewFileStore(cfgFileStore.Directory)
			case da.DaCelestia:
				cfgCelestia := da.NewCelestiaCfg()
				store, err = da.NewCelestiaStore(cfgCelestia)
				if err != nil {
					return err
				}
			default:
				return errors.New("unknown DA")
			}
			server := plasma.NewDAServer(cfgApp, store, slog.Default())
			server.Start()
			return nil
		},
	}
	startCmd.Flags().String("home", ".plasma-hub/config", "config file")

	return startCmd
}
