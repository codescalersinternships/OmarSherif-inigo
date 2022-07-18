package main

import (
	"errors"
	"os"
)

func LoadFromFile(path string) (*os.File, error) {
	if path == "" {
		return nil, errors.New("File path is empty")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
