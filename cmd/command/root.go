package command

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "plasma-da",
	Short: "Plasma DA is a simple CLI tool.",
	Long:  `Plasma DA is a simple CLI tool to roll up data to multiple services.`,
}

func Execute() error {
	err := rootCommand.Execute()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCommand.AddCommand(VersionCmd())
	rootCommand.AddCommand(StartCmd())
	rootCommand.AddCommand(InitConfigCmd())
}
