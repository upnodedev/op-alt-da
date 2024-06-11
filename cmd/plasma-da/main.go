package main

import (
	"fmt"
	"os"
	"plasma/cmd/command"
)

func main() {
	if err := command.Execute(); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		os.Exit(1)
	}
}
