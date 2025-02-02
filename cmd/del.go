package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "del",
	Short: "Remove file or folder",
	Long: `Remove file or folder from Directory

Example: 
	fts del -f -d <parent_directory> -n <file_name> -t <file_type>
	fts del -d <parent_directory> -n <folder_name>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if is_file {
			file_type = strings.TrimPrefix(file_type, ".")
			err := rootFolder.RemoveFile(dir, file_type, name)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Deleted file %s.%s from %s directory\n", name, file_type, dir)
			return
		}
		err := rootFolder.RemoveFolder(dir, name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Deleted folder %sfrom %s directory\n", name, dir)
	},
}
