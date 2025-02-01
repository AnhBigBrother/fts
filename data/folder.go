package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Folder struct {
	Directory  string    `json:"directory"`
	FolderName string    `json:"folder_name"`
	Files      []*File   `json:"files"`
	SubFolders []*Folder `json:"sub_folders"`
}

func NewFolder(dir, name string) Folder {
	return Folder{
		Directory:  dir,
		FolderName: name,
		Files:      []*File{},
		SubFolders: []*Folder{},
	}
}

func (root *Folder) validateFolder() error {
	if root.FolderName == "" {
		return errors.New("folder name cannot be empty")
	}
	return nil
}

func (root *Folder) FindFolderByDirectory(dir string) *Folder {
	if dir == "/" {
		return root
	}

	dir = strings.TrimPrefix(dir, "/")
	dir = strings.TrimSuffix(dir, "/")
	path := strings.Split(dir, "/")

	var curr *Folder = nil
	for _, p := range path {
		for _, fol := range root.SubFolders {
			if fol.FolderName == p {
				root = fol
				curr = fol
				break
			}
		}
	}

	return curr
}

func (root *Folder) AddFile(dir, file_type, name, content string) error {
	currentFolder := root.FindFolderByDirectory(dir)
	if currentFolder == nil {
		return fmt.Errorf("folder not found: %s", dir)
	}

	for _, f := range currentFolder.Files {
		if f.FileName == name && f.FileType == file_type {
			return fmt.Errorf("file %s.%s already existed", name, file_type)
		}
	}

	dir = strings.TrimPrefix(dir, "/")
	dir = strings.TrimSuffix(dir, "/")
	file_type = strings.TrimPrefix(file_type, ".")
	newFile := NewFile(fmt.Sprintf("/%s", dir), file_type, name, content)
	err := newFile.validateFile()
	if err != nil {
		return err
	}

	currentFolder.Files = append(currentFolder.Files, &newFile)
	return nil
}

func (root *Folder) RemoveFile(dir, file_type, file_name string) error {
	currentFolder := root.FindFolderByDirectory(dir)
	if currentFolder == nil {
		return fmt.Errorf("folder not found: %s", dir)
	}

	idx := -1
	for i, file := range currentFolder.Files {
		if file.FileName == file_name && file.FileType == file_type {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("no file '%s' in folder '%s'", file_name, dir)
	}

	currentFolder.Files = append(currentFolder.Files[:idx], currentFolder.Files[idx+1:]...)
	return nil
}

func (root *Folder) AddFolder(dir, name string) error {
	currentFolder := root.FindFolderByDirectory(dir)
	if currentFolder == nil {
		return fmt.Errorf("folder not found: %s", dir)
	}

	for _, f := range currentFolder.SubFolders {
		if f.FolderName == name {
			return fmt.Errorf("folder %s already existed", name)
		}
	}

	dir = strings.TrimPrefix(dir, "/")
	dir = strings.TrimSuffix(dir, "/")
	newFolder := NewFolder(fmt.Sprintf("/%s", dir), name)
	err := newFolder.validateFolder()
	if err != nil {
		return err
	}

	currentFolder.SubFolders = append(currentFolder.SubFolders, &newFolder)
	return nil
}

func (root *Folder) RemoveFolder(dir, folder_name string) error {
	currentFolder := root.FindFolderByDirectory(dir)
	if currentFolder == nil {
		return fmt.Errorf("folder not found: %s", dir)
	}

	idx := -1
	for i, f := range currentFolder.SubFolders {
		if f.FolderName == folder_name {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("no folder '%s' in folder '%s'", folder_name, dir)
	}

	currentFolder.SubFolders = append(currentFolder.SubFolders[:idx], currentFolder.SubFolders[idx+1:]...)
	return nil
}

func (root *Folder) RetrieveFolder(dir string) *Folder {
	return root.FindFolderByDirectory(dir)
}

func (root *Folder) RetrieveFile(dir, file_type, file_name string) *File {
	currentFolder := root.FindFolderByDirectory(dir)
	if currentFolder == nil {
		return nil
	}

	for _, f := range currentFolder.Files {
		if f.FileName == file_name && f.FileType == file_type {
			return f
		}
	}

	return nil
}

func (root *Folder) SearchFolderByName(folder_name string) []*Folder {
	ans := []*Folder{}

	var dfs func(root *Folder)
	dfs = func(root *Folder) {
		for _, f := range root.SubFolders {
			if f.FolderName == folder_name {
				ans = append(ans, f)
			}
			dfs(f)
		}
	}
	dfs(root)

	return ans
}

func (root *Folder) SearchFileByName(file_name string) []*File {
	ans := []*File{}

	var dfs func(root *Folder)
	dfs = func(root *Folder) {
		for _, file := range root.Files {
			if file.FileName == file_name {
				ans = append(ans, file)
			}
		}
		for _, folder := range root.SubFolders {
			dfs(folder)
		}
	}
	dfs(root)

	return ans
}

func (root *Folder) DisplayFolder() {
	folder := struct {
		Directory  string   `json:"directory"`
		FolderName string   `json:"folder_name"`
		Files      []string `json:"files"`
		SubFolders []string `json:"sub_folders"`
	}{
		Directory:  root.Directory,
		FolderName: root.FolderName,
		Files:      []string{},
		SubFolders: []string{},
	}
	for _, file := range root.Files {
		folder.Files = append(folder.Files, fmt.Sprintf("%s.%s", file.FileName, file.FileType))
	}
	for _, subFolder := range root.SubFolders {
		folder.SubFolders = append(folder.SubFolders, subFolder.FolderName)
	}

	data, err := json.MarshalIndent(folder, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func (root *Folder) DisplayFolderTree() {
	var dfs func(root *Folder, indent string)
	dfs = func(root *Folder, indent string) {
		fmt.Printf("%s%s\n", indent, root.FolderName)
		indent = fmt.Sprintf("   %s", indent)
		if len(root.SubFolders) > 0 {
			for _, subFolder := range root.SubFolders {
				dfs(subFolder, indent)
			}
		} else {
			fmt.Println()
		}
	}
	dfs(root, "|__ ")
}
