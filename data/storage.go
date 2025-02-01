package data

import (
	"encoding/json"
	"os"
	"strings"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](file_name string) *Storage[T] {
	return &Storage[T]{FileName: file_name}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			_, err := os.Create(s.FileName)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return json.Unmarshal(fileData, data)
}
