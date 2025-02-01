package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Add new file or folder",
	Long:  `Add new file or folder to directory`,
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
