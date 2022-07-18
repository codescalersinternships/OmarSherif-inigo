package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("not found")

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return definition, nil
}

func (d Dictionary) Append(key, value string) {
	d[key] = value
}

func (d Dictionary) AppendSection(key, value string) {
	d[key] = value
}
