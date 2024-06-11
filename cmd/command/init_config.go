package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

const DefaultHomeDir = ".plasma-da"

const ConfigTemplate = `[server]
http_host = "localhost"
http_port = 3128
da = "{{.Da}}"

[celestia]
rpc_port = "http://localhost:7980"
auth_token = ""
namespace = ""
max_block_size = 2000
gas_price = 0.002
eth_fallback_disabled = false

[filestore]
path = ".plasma-da/data/filestore"
`

func InitConfigCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir := cmd.Flag("home").Value.String()
			network := cmd.Flag("network").Value.String()
			da := cmd.Flag("da").Value.String()
			userDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			if homeDir == "" {
				homeDir = userDir + "/" + DefaultHomeDir
			}

			// validate da supported. Only file and celestia are supported
			if da != "file" && da != "celestia" {
				return fmt.Errorf("da %s is not supported", da)
			}

			// create the config file
			if err := createConfig(homeDir, network, da); err != nil {
				return err
			}

			return nil
		},
	}

	// set the flags
	initCmd.Flags().String("home", "", "config file of the plasma-da (default is $HOME/.plasma-da)")
	initCmd.Flags().String("network", "local", "network type")
	initCmd.Flags().String("da", "file", "data availability layer type")

	return initCmd
}

func createConfig(homeDir, network, da string) error {
	configPath := homeDir + "/config"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configPath, 0755); err != nil {
			return err
		}
	}

	data := struct {
		Da string
	}{
		Da: da,
	}
	t := template.Must(template.New("config").Parse(ConfigTemplate))

	// generate file config from template
	configFile := configPath + "/config.toml"
	fmt.Println("Creating config file at", configFile)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		outputFile, err := os.Create(configFile)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		err = t.Execute(outputFile, data)
		if err != nil {
			return err
		}

	} else {
		var input string
		fmt.Println("Config file already exists. Do you want to overwrite it? (y/n)")
		_, err := fmt.Scanln(&input)
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		if input == "yes" || input == "y" {
			fmt.Println("Overwriting config file....")
			outputFile, err := os.Create(configFile)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			err = t.Execute(outputFile, data)
			if err != nil {
				return err
			}
		}

		fmt.Println(`Config plasma-da is already initialized. Please check the config file at: `, configFile)
	}
	return nil
}
