package main

import (
	"alt-da/cmd/command"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const PrefixEnv = "ALT_DA"

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(PrefixEnv)
	if err := command.Execute(); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		os.Exit(1)
	}
}
