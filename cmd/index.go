package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

var IndexPathParameter string
var BaseDirectoryParameter string

var createIndexCmd = &cobra.Command{
	Use: "index",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Index file: %s\n", IndexPathParameter)
		fmt.Printf("Base directory: %s\n", BaseDirectoryParameter)

		docxFilePattern := regexp.MustCompile(`.*\.docx`)
		err := filepath.Walk(BaseDirectoryParameter,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				match := docxFilePattern.FindStringSubmatch(path)
				if len(match) == 1 {
					fmt.Println(path)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	createIndexCmd.Flags().StringVarP(&IndexPathParameter, "index-file", "i", "", "path of the generated index file")
	createIndexCmd.Flags().StringVarP(&BaseDirectoryParameter, "base-directory", "b", "", "directory where to start to search file docx files")
	if err := createIndexCmd.MarkFlagRequired("index-file"); err != nil {
		panic(err)
	}
	if err := createIndexCmd.MarkFlagRequired("base-directory"); err != nil {
		panic(err)
	}
}
