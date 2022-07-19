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
		p.allSections.PutKey(section.name, section)
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
			return SyntaxError // if the section open but not closed
		} else if strings.Contains(statment, "]") && !strings.Contains(statment, "[") {
			return SyntaxError // if the section is closed but not opened
		} else if (strings.Contains(statment, "[") && strings.Contains(statment, "]")) && (openBracket > closeBracket || strings.Count(statment, "[") != 1 || strings.Count(statment, "]") != 1) {
			return SyntaxError // if the section is closed before opened or more than one section
		} else if strings.Contains(statment, ";") && string(statment[0]) != ";" {
			return SyntaxError // if the statment starts contains ; but doest not start with ;
		} else if strings.Contains(statment, "=") && string(statment[0]) == "=" {
			return SyntaxError // if the first character is =
		} else if (strings.Contains(statment, "]") && strings.Contains(statment, "[")) && (openBracket > closeBracket) {
			return SyntaxError // if the brackets are not in the right order
		} else if !strings.Contains(statment, "]") && !strings.Contains(statment, "[") && !strings.Contains(statment, "=") && !strings.Contains(statment, ";") {
			return SyntaxError // if the statment is not a comment and does not contain brackets
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
		return sections, NoSectionError
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
	return section, NoSectionError
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
		p.allSections.PutKey(section_name, section)
	}
	section.SetKey(key, value)

}

// returns the section with the given name
func (p *Parser) LoadFromFile(fileName string) error {
	// we read the file
	fileContent, err := ReadFile(fileName)
	if err != nil {
		return err
	}
	err = p.LoadFromString(string(fileContent))
	return err
}

func (p *Parser) SaveToFile(fileName string) error {
	err := WriteToFile(fileName, fmt.Sprintf("Map: %v", p.allSections))
	return err
}
