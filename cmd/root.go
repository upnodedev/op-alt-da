package main

import (
	"github.com/spf13/cobra"
	"plasma/cmd/command"
)

var rootCommand = &cobra.Command{
	Use:   "plasma-da",
	Short: "Plasma Hub is a simple CLI tool.",
	Long:  `Plasma Hub is a simple CLI tool to roll up data to multiple services.`,
}

func Execute() error {
	err := rootCommand.Execute()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCommand.AddCommand(command.VersionCmd())
	rootCommand.AddCommand(command.StartCmd())
	rootCommand.AddCommand(command.InitConfigCmd())
}
