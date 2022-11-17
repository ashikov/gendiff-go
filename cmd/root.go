package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var formatter = "stylish"
var rootCmd = &cobra.Command{
	Use:   "gendiff",
	Short: "Application that finds differences in configuration files",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("formatter", formatter, "Set output formatter")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
