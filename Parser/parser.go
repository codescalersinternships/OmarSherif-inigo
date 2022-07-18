package main

import (
	"errors"
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
func (p *Parser) LoadFromString(input string) (string, error) {
	p.code = input
	// we seperat the sections
	sections := strings.Split(input, "[")
	// we loop over across the sections
	for _, section := range sections {
		// we create a section object individually
		section, err := SectionConstructor(section)
		if err != nil {
			return "", err
		}
		// we append the section to the parser sections
		p.allSections.Append(section.name, section)
	}
	return "", nil
}

// returns the all the sections
func (p *Parser) GetSections() (sectionDictionary, error) {
	return p.allSections, nil
}

// returns the all sections names
func (p *Parser) GetSectionNames() ([]string, error) {
	sections := []string{}
	for sectionName, _ := range p.allSections.sections {
		sections = append(sections, sectionName)
	}
	if len(sections) == 0 {
		return sections, NoSectionError
	}
	return sections, nil
}

// returns the section with the given name
func (p *Parser) GetSection(sectionName string) (Section, error) {
	section := Section{}
	for sectionName, _ := range p.allSections.sections {
		if sectionName == sectionName {
			section = p.allSections.sections[sectionName]
			return section, nil
		}
	}
	// Didnt find the section
	return section, NoSectionError
}
