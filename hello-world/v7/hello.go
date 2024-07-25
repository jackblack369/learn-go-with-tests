package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const china = "China"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const chnHelloPrefix = "你好, "

// Hello returns a personalised greeting in a given language.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case china:
		prefix = chnHelloPrefix

	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "China"))
}
