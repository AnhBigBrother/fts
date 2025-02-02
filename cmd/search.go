package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var search = &cobra.Command{
	Use:   "search",
	Short: "Search file/folder by name",
	Long: `Search file/folder by name 

Example: 
	fts search -f <file_name>
	fts search <folder_name> 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		search_name := strings.Join(args, " ")
		if is_file {
			files := rootFolder.SearchFileByName(search_name)
			fmt.Printf("found %d files named '%s':\n", len(files), search_name)
			for _, f := range files {
				f.DisplayFile()
			}
			return
		}
		folders := rootFolder.SearchFolderByName(search_name)
		fmt.Printf("found %d folders named '%s':\n", len(folders), search_name)
		for _, f := range folders {
			f.DisplayFolder()
		}
	},
}
