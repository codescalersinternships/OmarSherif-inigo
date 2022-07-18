package main

import "fmt"

const spanish = "spanish"
const english = "english"
const french = "french"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return GreetingPrefix(language) + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
func main() {
	fmt.Println(Hello("mom", "french"))
}
