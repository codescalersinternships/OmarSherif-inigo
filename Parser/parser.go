package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	commentStart      = ';'
	sectionStart      = '['
	sectionEnd        = ']'
	keyValueSeperator = '='
	noKeyMessage      = "Key %s is not found in the following section  %s"
	noSectionMessage  = "Section %s is not found"
)

// Errors
var (
	NoKeyError     = errors.New(noKeyMessage)
	NoKeysError    = errors.New("There is no keys %s in section named %s")
	NoSectionError = errors.New(noSectionMessage)
	NoSections     = errors.New("There is no sections")
	SyntaxError    = errors.New("There is no section named %s")
)

type Parser struct {
	code        string
	allSections sectionDictionary
}

func NewParser() *Parser {
	return &Parser{"1", sectionDictionary{make(map[string]Section)}}
}
func (p *Parser) LoadFromString(input string) {
	p.code = input
	// we seperat the sections
	sections := strings.Split(input, "[")
	// we loop over across the sections
	for index, section := range sections {
		// we create a section object individually

		if index == 0 {
			continue
		}
		section, err := NewSection(section)
		if err != nil {

		}
		// we append the section to the parser sections
		p.allSections.Append(section.name, section)
	}

}

// returns the all the sections
func (p *Parser) GetSections() (sectionDictionary, error) {
	return p.allSections, nil
}

// returns the all sections names
func (p *Parser) GetSectionNames() ([]string, error) {
	sections := []string{}
	for sectionName := range p.allSections.sections {
		sections = append(sections, sectionName)
	}
	if len(sections) == 0 {
		return sections, NoSectionError
	}
	return sections, nil
}

// returns the section with the given name
func (p *Parser) GetSection(sectionName string) (Section, error) {
	var section Section
	for name, _ := range p.allSections.sections {
		if sectionName == name {
			section = p.allSections.sections[sectionName]
			return section, nil
		}
	}

	// Didnt find the section
	return section, NoSectionError
}

// returns the section's key value with the given section and key names
func (p *Parser) GetValue(sectionName string, key string) (string, error) {
	section, err := p.GetSection(sectionName)
	if err != nil {
		return "", err
	}
	return section.GetValue(key)
}

// returns the section's comment's list with the given section name
func (p *Parser) GetComments(sectionName string) ([]string, error) {
	section, err := p.GetSection(sectionName)
	if err != nil {
		return nil, err
	}
	return section.GetComments(), nil
}

// it returns it into its original Form
func (p *Parser) ToString(input string) (output string) {
	for sectionName, section := range p.allSections.sections {
		output += fmt.Sprintf("[%s]\n", sectionName)
		for key, value := range section.dictionary {
			output += fmt.Sprintf("%s = %s\n", key, value)
		}
		for _, comment := range section.commentList {
			output += fmt.Sprintf(";%s\n", comment)
		}
		output += "\n"

	}

	return output
}

// returns the section with the given name
func (p *Parser) LoadFromFile(fileName string) {

	// we read the file
	fileContent := ReadFile(fileName)
	p.LoadFromString(string(fileContent))
}

func (p *Parser) SaveToFile(fileName string) {
	WriteToFile(fileName, fmt.Sprintf("Map: %v", p.allSections))
}
