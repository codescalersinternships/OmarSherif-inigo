package main

import (
	"errors"
	"fmt"
	"strings"
)

// Errors
var (
	ErrNoSection   = errors.New("Section %s is not found")
	ErrNoSections  = errors.New("There is no sections")
	ErrSyntax      = errors.New("There is no section named %s")
	ErrKeyNotFound = errors.New("There is no keys %s in section named %s")
)

type sectionDictionary struct {
	sections map[string]Section
}

type Parser struct {
	allSections sectionDictionary
}

func NewParser() *Parser {
	return &Parser{sectionDictionary{make(map[string]Section)}}
}
func (p *Parser) LoadFromString(input string) error {
	// we check the input if it have Syntax Errors
	err := p.checkInput(input)
	if err != nil {
		return err
	}
	// we seperate the sections
	sections := strings.Split(input, "[")
	// we loop over across the sections
	for index, section := range sections {
		// we create a section object individually

		if index == 0 {
			continue
		}
		section, err := NewSection(section)
		if err != nil {
			return err
		}

		// we append the section to the parser sections
		p.allSections.sections[section.name] = section

	}
	return nil
}
func (p *Parser) checkInput(input string) error {
	statements := strings.Split(input, "\n")
	for _, statment := range statements {
		statment = strings.TrimSpace(statment)
		openBracket := strings.Index(statment, "[")
		closeBracket := strings.Index(statment, "]")
		if len(statment) == 0 {
			continue
		}
		if strings.Contains(statment, "[") && !strings.Contains(statment, "]") {
			return ErrSyntax // if the section open but not closed
		} else if strings.Contains(statment, "]") && !strings.Contains(statment, "[") {
			return ErrSyntax // if the section is closed but not opened
		} else if (strings.Contains(statment, "[") && strings.Contains(statment, "]")) && (openBracket > closeBracket || strings.Count(statment, "[") != 1 || strings.Count(statment, "]") != 1) {
			return ErrSyntax // if the section is closed before opened or more than one section
		} else if strings.Contains(statment, ";") && string(statment[0]) != ";" {
			return ErrSyntax // if the statment starts contains ; but doest not start with ;
		} else if strings.Contains(statment, "=") && (string(statment[0]) == "=" ) {
			return ErrSyntax // if the first character is =
		} else if strings.Contains(statment, "]") && strings.Contains(statment, "[") && len(statment) == 2 {
			return ErrSyntax // if the first character is not =
		} else if (strings.Contains(statment, "]") && strings.Contains(statment, "[")) && (openBracket > closeBracket) {
			return ErrSyntax // if the brackets are not in the right order
		} else if !strings.Contains(statment, "]") && !strings.Contains(statment, "[") && !strings.Contains(statment, "=") && !strings.Contains(statment, ";") {
			return ErrSyntax // if the statment is not a comment and does not contain brackets
		}

	}
	return nil

}

// returns the all the sections (TESTED)
func (p *Parser) GetSections() (sectionDictionary, error) {
	return p.allSections, nil
}

// returns the all sections names (TESTED)
func (p *Parser) GetSectionNames() ([]string, error) {
	sections := []string{}
	for sectionName := range p.allSections.sections {
		sections = append(sections, sectionName)
	}
	if len(sections) == 0 {
		return sections, ErrNoSection
	}
	return sections, nil
}

// returns the section with the given name (TESTED)
func (p *Parser) GetSection(sectionName string) (Section, error) {
	var section Section
	for name, _ := range p.allSections.sections {
		if sectionName == name {
			section = p.allSections.sections[sectionName]
			return section, nil
		}
	}

	// Didnt find the section
	return section, ErrNoSection
}

// returns the section's key value with the given section and key names (TESTED)
func (p *Parser) Get(sectionName string, key string) (string, error) {
	section, err := p.GetSection(sectionName)
	if err != nil {
		return "", err
	}
	return section.GetValue(key)
}

// it returns it into its original Form
func (p Parser) String(input string) (output string) {
	for sectionName, section := range p.allSections.sections {
		output += fmt.Sprintf("[%s]\n", sectionName)
		for key, value := range section.dictionary {
			output += fmt.Sprintf("%s = %s\n", key, value)
		}
		output += "\n"
	}
	return output
}

// set a key in a section
func (p *Parser) Set(section_name, key, value string) {
	section, err := p.GetSection(section_name)
	if err != nil { // if the section is not found
		section, _ = NewSection(section_name)
		p.allSections.sections[section.name] = section
	}
	section.SetKey(key, value)
}

// loads a file into the parser
func (p *Parser) LoadFromFile(fileName string) error {
	// we read the file
	fileContent, err := ReadFile(fileName)
	if err != nil {
		return err
	}
	err = p.LoadFromString(string(fileContent))
	return err
}

// save the sections into the file
func (p *Parser) SaveToFile(fileName string) error {
	err := WriteToFile(fileName, fmt.Sprintf("Map: %v", p.allSections))
	return err
}
