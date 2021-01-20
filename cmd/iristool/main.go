package main

import (
	"os"

	"github.com/dreamer-zq/diff/cmd/iristool/cmd"
)

func main() {
	rootCmd := cmd.NewToolCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
