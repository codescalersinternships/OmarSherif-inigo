package main

import (
	"fmt"
)

const input = `[owner]
	name = John Doe
	organization = Acme Widgets Inc.

	[database]
	; use IP address in case network name resolution is not working
	server = 192.0.2.62
	port = 143
	file = "payroll.dat"`

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
