package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var search = &cobra.Command{
	Use:   "search",
	Short: "search file/folder by name",
	Long:  `search file/folder by name`,
	Run: func(cmd *cobra.Command, args []string) {
		if is_file {
			files := rootFolder.SearchFileByName(name)
			fmt.Printf("found %d files named %s:\n", len(files), name)
			for _, f := range files {
				f.DisplayFile()
			}
			return
		}
		folders := rootFolder.SearchFolderByName(name)
		fmt.Printf("found %d folders named %s:\n", len(folders), name)
		for _, f := range folders {
			f.DisplayFolder()
		}
	},
}
