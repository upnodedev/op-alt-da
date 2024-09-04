package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"plasma/version"
)

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of alt-da",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("💈 Alt DA version", version.BuildVersion)
			fmt.Println("💈 Build time:", version.BuildTime)
			fmt.Println("💈 Git commit:", version.BuildCommit)
		},
	}
}
