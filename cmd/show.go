package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var show = &cobra.Command{
	Use:   "show",
	Short: "Show folder-tree structure",
	Long: `Display folder structure as a tree
	
Example: 
	fts show -d <folder_directory>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		folder := rootFolder.FindFolderByDirectory(dir)
		if folder == nil {
			fmt.Println("folder not found:", folder)
			return
		}
		folder.DisplayFolderTree()
	},
}
