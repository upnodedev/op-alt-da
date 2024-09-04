package command

import (
	opaltda "alt-da"
	"alt-da/config"
	"alt-da/da"
	"alt-da/da/arweave"
	"alt-da/da/celestia"
	"alt-da/da/file"
	"alt-da/da/ipfs"
	"alt-da/evm"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"os"
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
				homeDir = userDir + "/.alt-da"
			}
			submitter, err := evm.NewSubmitter(cfgApp)
			if err != nil {
				return err
			}
			var daId [32]byte
			copy(daId[:], cfgApp.DaID)

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
				store, err = celestia.NewCelestiaStore(cfgCelestia, daId, submitter)
				if err != nil {
					return err
				}
			case ipfs.DaIpfs:
				cfgIpfs := ipfs.ParseConfig(cmd)
				store, err = ipfs.NewIpfsStore(cfgIpfs, daId, submitter)
				if err != nil {
					return err
				}
			case arweave.DaAr:
				cfgAr := arweave.ParseConfig(cmd)
				store, err = arweave.NewArStore(cfgAr, daId, submitter)
				if err != nil {
					return err
				}
			default:
				return errors.New(fmt.Sprintf("unknown DA type: %s", cfgApp.DA))
			}
			server := opaltda.NewDAServer(cfgApp, store, slog.Default())
			server.Start()
			return nil
		},
	}
	AppFlags(startCmd)
	celestia.AddFlags(startCmd)
	file.AddFlags(startCmd)
	ipfs.AddFlags(startCmd)
	arweave.AddFlags(startCmd)

	return startCmd
}

func AppFlags(cmd *cobra.Command) {
	cmd.Flags().String("host", "", "host (default is localhost)")
	cmd.Flags().Int("port", 3128, "port (default is 3128)")
	cmd.Flags().String("home", "", "config file (default is $HOME/.alt-da)")
	cmd.Flags().String("da", "", "data availability layer type (default is file store)")
	cmd.Flags().String("da-id", "", "data availability layer id (default is 0x000c for celestia)")
	cmd.Flags().String("evm-rpc-url", "", "the rpc url for the evm")
	cmd.Flags().String("key-file", "", "the key file of account use for submitting data mapping")
	cmd.Flags().String("passphrase", "", "the passphrase for the key file")
	cmd.Flags().Int64("chain-id", 0, "the chain id for the evm")
	cmd.Flags().String("alt-da-hub-addr", "", "the alt da hub address")
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
	if daId := cmd.Flag("da-id").Value.String(); daId != "" {
		cfgApp.DaID = daId
	}
	if evmRpcUrl := cmd.Flag("evm-rpc-url").Value.String(); evmRpcUrl != "" {
		cfgApp.EvmRpcUrl = evmRpcUrl
	}
	if keyFile := cmd.Flag("key-file").Value.String(); keyFile != "" {
		cfgApp.KeyFile = keyFile
	}
	if passphrase := cmd.Flag("passphrase").Value.String(); passphrase != "" {
		cfgApp.Passphrase = passphrase
	}
	if chainId, err := cmd.Flags().GetInt64("chain-id"); err == nil {
		cfgApp.ChainId = chainId
	}
	if hubAddr := cmd.Flag("alt-da-hub-addr").Value.String(); hubAddr != "" {
		cfgApp.AltDaHubAddr = hubAddr
	}

	return cfgApp
}
