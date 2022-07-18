package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}

func (d Dictionary) Append(key, value string) {
	d[key] = value
}

func (d Dictionary) AppendSection(key, value string) {
	d[key] = value
}
