package main

import "fmt"

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
	/* 	sections, err := parser.GetSections()
	   	if err != nil {
	   		panic(err)
	   	}
	   	for _, section := range sections {
	   		fmt.Printf("Section: %s\n", section.Name())
	   		for _, key := range section.GetKeyList() {
	   			fmt.Printf("\t%s = %s\n", key, section.Get(key))
	   		}
	   	} */

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
