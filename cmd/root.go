package cmd

import (
	"fmt"
	"os"

	"github.com/AnhBigBrother/fts/data"
	"github.com/spf13/cobra"
)

var (
	rootFolder data.Folder
	store      *data.Storage[data.Folder]
	rootCmd    *cobra.Command

	is_file   bool
	dir       string
	name      string
	file_type string
)

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
	rootFolder = data.Folder{
		Directory:  "",
		FolderName: "/",
		Files:      []*data.File{},
		SubFolders: []*data.Folder{},
	}

	homeDir, _ := os.UserHomeDir()
	store = data.NewStorage[data.Folder](fmt.Sprintf("%s/fts_data.json", homeDir))

	store.Load(&rootFolder)

	rootCmd = &cobra.Command{
		Use:   "fts",
		Short: "Folder tree structure",
		Long:  "A CLI application that simulates a folder-tree structure",
	}

	rootCmd.AddCommand(add)
	rootCmd.AddCommand(get)
	rootCmd.AddCommand(delete)
	rootCmd.AddCommand(search)
	rootCmd.AddCommand(show)

	rootCmd.PersistentFlags().BoolVarP(&is_file, "file", "f", false, "set is file or not")
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "/", "file/forder's parent directory")
	rootCmd.PersistentFlags().StringVarP(&file_type, "type", "t", "txt", "file type")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "NAME", "file/folder name")
}
