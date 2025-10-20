package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "moti",
	Short: "a cross-platform CLI app for searching text in DOCX documents",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing moti")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
