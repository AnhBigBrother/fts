package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "del",
	Short: "Remove file or folder",
	Long:  `Remove file or folder from directory`,
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
