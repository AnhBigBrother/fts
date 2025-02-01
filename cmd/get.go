package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "Retrieve file/folder",
	Long:  `Retrieve file/folder from directory`,
	Run: func(cmd *cobra.Command, args []string) {
		if is_file {
			file_type = strings.TrimPrefix(file_type, ".")
			file := rootFolder.RetrieveFile(dir, file_type, name)
			if file == nil {
				fmt.Println("file not found")
				return
			}
			file.DisplayFile()
			return
		}
		folder := rootFolder.RetrieveFolder(dir)
		if folder == nil {
			fmt.Println("folder not found")
			return
		}
		folder.DisplayFolder()
	},
}
