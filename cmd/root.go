package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"errors"

	"github.com/spf13/cobra"
)

func isFileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	 } else {
		message := fmt.Sprintf("file %s does not exist", path)
		return false, errors.New(message)
	 }
}

var rootCmd = &cobra.Command{
	Use:   "gendiff path1 path2",
	Short: "Application that finds differences in configuration files",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		formatter, err := cmd.Flags().GetString("formatter")
		if err !=nil {
			log.Fatal(err)
		}

		path1, path2 := args[0], args[1]

		absPath1, err := filepath.Abs(path1)
		if err != nil {
			log.Fatal(err)
		}

		absPath2, err := filepath.Abs(path2)
		if err != nil {
			log.Fatal(err)
		}

		_, err = isFileExists(absPath1)
		if err != nil {
			log.Fatal(err)
		}

		_, err = isFileExists(absPath2)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Yo!", formatter)
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
