package cmd

import (
	"fmt"
	"os"

	"github.com/AnhBigBrother/folder-tree-structure/data"
	"github.com/spf13/cobra"
)

var (
	is_file   bool
	dir       string
	name      string
	file_type string
)

var rootFolder = data.Folder{
	Directory:  "",
	FolderName: "/",
	Files:      []*data.File{},
	SubFolders: []*data.Folder{},
}
var store = data.NewStorage[data.Folder]("./fts_data.json")

var rootCmd = &cobra.Command{
	Use:   "fts",
	Short: "Folder tree structure",
	Long:  "A folder tree structure CLI application following the object-oriented programming principles",
}

func Excute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = store.Save(rootFolder)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	store.Load(&rootFolder)

	rootCmd.AddCommand(add)
	rootCmd.AddCommand(get)
	rootCmd.AddCommand(delete)
	rootCmd.AddCommand(search)
	rootCmd.AddCommand(tree)

	rootCmd.PersistentFlags().BoolVar(&is_file, "file", false, "set is file or not")
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "/", "file/forder directory")
	rootCmd.PersistentFlags().StringVarP(&file_type, "type", "t", "txt", "file type")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "NAME", "file/folder name")
}
