package main

import (
	"context"
	"os"

	"playground/modules"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "playground",
		Short: "Go playground for experimenting with various modules",
	}

	rootCmd.AddCommand(modules.NewSequencesCmd())

	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}
