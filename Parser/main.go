package main

import (
	"fmt"
)

/*
func main() {
	const input = `[owner]
	name = John Doe
	organization = Acme Widgets Inc.

	[database]
	; use IP address in case network name resolution is not working
	server = 192.0.2.62
	port = 143
	file = "payroll.dat"`

	parser := NewParser()
	parser.LoadFromString(input)

	fmt.Println("\n\n")
	fmt.Println("Section names:")
	sectionNames, err := parser.GetSectionNames()
	if err != nil {
		panic(err)
	}
	for _, sectionName := range sectionNames {
		fmt.Printf("\t%s\n", sectionName)
	}

}
*/

func main() {
	fmt.Println("Enter filename: ")
	filename := "omar.txt"
	input := "I love pizza"

	// file is created and read
	CreateFile(filename, input)
	ReadFile(filename)
}
