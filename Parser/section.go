package main

import (
	"strings"
)

type Section struct {
	name        string
	commentList []string
	keyList     []string
	dictionary  Dictionary
}

func NewSection(input string) (Section, error) {

	sectionName, commentList, keyList, sectionDictionary, err := createSection(input)

	return Section{sectionName, commentList, keyList, sectionDictionary}, err
}

//
func createSection(input string) (string, []string, []string, Dictionary, error) {
	// Section Variables
	var sectionName string
	var commentList []string
	var keyList []string
	var sectionDictionary Dictionary
	var sectionDictionaryMap = make(map[string]string)

	// section
	sections := strings.Split(input, "]")

	sectionName = strings.TrimSpace(sections[0])

	// section  statements
	sections[1] = strings.TrimSpace(sections[1])
	sectionStatements := strings.Split(sections[1], "\n")

	for _, statement := range sectionStatements {
		// we remove the spaces
		statement = strings.TrimSpace(statement)
		// we check if the statement is a comment
		if strings.HasPrefix(statement, string(commentStart)) {
			commentList = append(commentList, strings.TrimSpace(statement))
		} else {
			// we split statement into a key value pair
			keyValue := strings.Split(statement, "=")
			keyList = append(keyList, strings.TrimSpace(keyValue[0]))
			sectionDictionaryMap[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
		}
	}
	sectionDictionary = Dictionary(sectionDictionaryMap)
	return sectionName, commentList, keyList, sectionDictionary, nil
}

// Name returns name of Section.
func (s *Section) Name() string {
	return s.name
}

// GetKeys returns all the keys in the section.
func (s *Section) GetKeyList() (keys []string) {
	return s.keyList
}

// GetValue returns the value of a key in the section.
func (s *Section) GetValue(key string) (string, error) {
	return s.dictionary.Search(key)
}

// Set a key in the section.
func (s *Section) SetKey(key string, value string) {
	s.dictionary.Append(key, value)
}

// Update the value of a key in the section.
func (s *Section) UpdateKey(key string, value string) error {
	_, err := s.dictionary.Search(key)
	if err != nil {
		return err
	} else {
		s.SetKey(key, value)
		return nil
	}
}

// returns the Comment List of the section
func (s *Section) GetComments() []string {
	return s.commentList
}