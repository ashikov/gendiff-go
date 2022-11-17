package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gendiff path1 path2",
	Short: "Application that finds differences in configuration files",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		formatter, _ := cmd.Flags().GetString("formatter")
		fmt.Println(formatter)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("formatter", "stylish", "Set output formatter")
}
