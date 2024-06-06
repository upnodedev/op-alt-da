package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"os"
	"plasma"
	"plasma/config"
	"plasma/da"
)

func StartCmd() *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			userDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			homeDir := cmd.Flag("home").Value.String()
			if homeDir == "" {
				homeDir = userDir + "/.plasma-da"
			}
			cfgApp := config.NewAppConfig(homeDir)

			var store da.KVStore
			switch cfgApp.DA {
			case da.DaFile:
				cfgFileStore := da.NewFileStoreCfg()
				store, err = da.NewFileStore(cfgFileStore.Directory)
				if err != nil {
					return err
				}
			case da.DaCelestia:
				cfgCelestia := da.NewCelestiaCfg()
				store, err = da.NewCelestiaStore(cfgCelestia, homeDir)
				if err != nil {
					return err
				}
			default:
				return errors.New(fmt.Sprintf("unknown DA type: %s", cfgApp.DA))
			}
			server := plasma.NewDAServer(cfgApp, store, slog.Default())
			server.Start()
			return nil
		},
	}
	startCmd.Flags().String("home", "", "config file (default is $HOME/.plasma-da)")

	return startCmd
}
