package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "TBD",
	Long:  "TBD",
}

var IndexPathParameter string
var BaseDirectoryParameter string

var createIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "TBD",
	Long:  "TBD",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Index file: %s\n", IndexPathParameter)
		fmt.Printf("Base directory: %s\n", BaseDirectoryParameter)
	},
}

func init() {
	createIndexCmd.Flags().StringVarP(&IndexPathParameter, "index-file", "i", "", "The path of the generated index file")
	createIndexCmd.Flags().StringVarP(&BaseDirectoryParameter, "base-directory", "b", "", "The directory where to start to search file docx files")
	if err := createIndexCmd.MarkFlagRequired("index-file"); err != nil {
		panic(err)
	}
	if err := createIndexCmd.MarkFlagRequired("base-directory"); err != nil {
		panic(err)
	}

	createCmd.AddCommand(createIndexCmd)
	rootCmd.AddCommand(createCmd)
}
