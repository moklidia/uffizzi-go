package cmd

import (
	"fmt"
	"os"

	"UffizziCloud/uffizzi-go/cmd/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uffizzi",
	Short: "A command-line interace (CLI) for Uffizzi App",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
}
