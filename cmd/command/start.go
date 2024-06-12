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
	"plasma/da/celestia"
	"plasma/da/file"
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
			cfgApp := ParseAppFlags(cmd)
			homeDir := cfgApp.HomeDir
			if homeDir == "" {
				homeDir = userDir + "/.plasma-da"
			}

			var store da.KVStore
			switch cfgApp.DA {
			case file.DaFile:
				cfgFileStore := file.ParseConfig(cmd)
				store, err = file.NewFileStore(cfgFileStore)
				if err != nil {
					return err
				}
			case celestia.DaCelestia:
				cfgCelestia := celestia.ParseConfig(cmd)
				store, err = celestia.NewCelestiaStore(cfgCelestia, homeDir)
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
	AppFlags(startCmd)
	celestia.AddFlags(startCmd)
	file.AddFlags(startCmd)

	return startCmd
}

func AppFlags(cmd *cobra.Command) {
	cmd.Flags().String("host", "", "host (default is localhost)")
	cmd.Flags().Int("port", 3128, "port (default is 3128)")
	cmd.Flags().String("home", "", "config file (default is $HOME/.plasma-da)")
	cmd.Flags().String("da", "", "data availability layer type (default is file store)")
}

func ParseAppFlags(cmd *cobra.Command) config.App {
	cfgApp := config.DefaultConfig()
	if host := cmd.Flag("host").Value.String(); host != "" {
		cfgApp.Host = host
	}
	if port, err := cmd.Flags().GetInt("port"); err == nil {
		cfgApp.Port = port
	}
	if homeDir := cmd.Flag("home").Value.String(); homeDir != "" {
		cfgApp.HomeDir = homeDir
	}
	if daLayer := cmd.Flag("da").Value.String(); daLayer != "" {
		cfgApp.DA = daLayer
	}

	return cfgApp
}
