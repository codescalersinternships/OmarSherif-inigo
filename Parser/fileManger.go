// Golang program to read and write the files
package main

// importing the requires packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filename, text string) {

	// fmt package implements formatted I/O and
	// contains inbuilt methods like Printf
	// and Scanf
	fmt.Printf("Writing to a file in Go lang\n")

	// Creating the file using Create() method
	// with user inputted filename and err
	// variable catches any error thrown
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// closing the running file after the main
	// method has completed execution and
	// the writing to the file is complete
	defer file.Close()

	// writing data to the file using
	// WriteString() method and the
	// length of the string is stored
	// in len variable
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile(filename string) string {

	// we open the file
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	// we read the file
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	// we close the file
	file.Close()
	// we load the file content into the parser
	return string(fileContent)

}

func WriteToFile(filename, text string) {

	// we open the file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	// we write the file
	_, err = file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	// we close the file
	file.Close()
}
