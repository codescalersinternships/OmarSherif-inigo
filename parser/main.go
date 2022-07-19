package main

import (
	"fmt"
)

func main() {

	parser := NewParser()
	parser.LoadFromString(input)

	fmt.Println("Section names:")

	// print array of section names
	sectionNames, err := parser.GetSectionNames()
	if err != nil {
		panic(err)
	}
	for _, sectionName := range sectionNames {
		fmt.Printf("\t%s\n", sectionName)
	}
	fmt.Println()

	fmt.Println("Section Keys:")

	fmt.Println(parser.GetSections())

}
