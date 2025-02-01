package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type File struct {
	Directory string `json:"directory"`
	FileName  string `json:"file_name"`
	Content   string `json:"content"`
	FileType  string `json:"file_type"`
}

func NewFile(dir, file_type, name, content string) File {
	return File{
		Directory: dir,
		FileName:  name,
		Content:   content,
		FileType:  file_type,
	}
}

func (file *File) validateFile() error {
	err_msg := []string{}
	if file.FileName == "" {
		err_msg = append(err_msg, "file name cannot be empty")
	}
	if file.FileType == "" {
		err_msg = append(err_msg, "file type is required")
	}
	if len(err_msg) == 0 {
		return nil
	}
	return errors.New(strings.Join(err_msg, ", "))
}

func (file *File) DisplayFile() {
	data, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
