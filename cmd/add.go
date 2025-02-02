package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Add new file or folder",
	Long: `Add new file or folder to directory
	
Example: 
	fts add -f -d <parent_directory> -n <file_name> -t <file_type> <file_content>
	fts add -d <parent_directory> -n <folder_name>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if is_file {
			file_type = strings.TrimPrefix(file_type, ".")
			err := rootFolder.AddFile(dir, file_type, name, strings.Join(args, " "))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("added file %s.%s to directory %s \n", name, file_type, dir)
			return
		}
		err := rootFolder.AddFolder(dir, name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("added folder %s to directory %s \n", name, dir)
	},
}
