package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"plasma/cmd/command"
)

const PrefixEnv = "PLASMA_DA"

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(PrefixEnv)
	if err := command.Execute(); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		os.Exit(1)
	}
}
