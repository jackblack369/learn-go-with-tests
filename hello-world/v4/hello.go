package main

import "fmt"

const englishHelloPrefix = "Hello, "
const my_name = "brook"

// Hello returns a personalised greeting.
func Hello(name string) string {
	return my_name + name
}

func main() {
	fmt.Println(Hello("world"))
}
