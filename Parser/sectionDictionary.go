package main

import "errors"

type sectionDictionary struct {
	sections map[string]Section
}

func (d sectionDictionary) Search(key string) (Section, error) {
	definition, ok := d.sections[key]
	if !ok {
		return Section{}, errors.New("could not find the word you were looking for")
	}

	return definition, nil
}

func (d sectionDictionary) Append(key string, value Section) {
	d.sections[key] = value
}
