package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tree = &cobra.Command{
	Use:   "tree",
	Short: "Show folder structure tree",
	Long:  `Display folder structure as a tree`,
	Run: func(cmd *cobra.Command, args []string) {
		folder := rootFolder.FindFolderByDirectory(dir)
		if folder == nil {
			fmt.Println("folder not found:", folder)
			return
		}
		folder.DisplayFolderTree()
	},
}
