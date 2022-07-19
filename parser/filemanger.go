// Golang program to read and write the files
package main

// importing the requires packages
import (
	"io/ioutil"
	"os"
)

func CreateFile(filename, text string) (int, error) {

	// Creating the file using Create() method
	// with user inputted filename and err
	// variable catches any error thrown
	file, err := os.Create(filename)

	if err != nil {
		return 0, err
	}
	defer file.Close()

	// closing the running file after the main
	// method has completed execution and
	// the writing to the file is complete

	// writing data to the file using
	// WriteString() method and the
	// length of the string is stored
	// in len variable
	len, err := file.WriteString(text)
	if err != nil {
		return 0, err
	}

	return len, err

}

func ReadFile(filename string) (string, error) {

	// we open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	// we read the file
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	// we close the file
	file.Close()
	// we load the file content into the parser
	return string(fileContent), err

}

func WriteToFile(filename, text string) error {

	// we open the file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// we write the file
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	// we close the file
	file.Close()
	return nil
}
