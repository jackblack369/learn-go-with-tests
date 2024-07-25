package main

import "fmt"

// Hello returns a greeting.
func Hello() string {
	return "Hello, world"
}

func HelloBrook() string {
	return "Hello, Brook"
}

func main() {
	fmt.Println(Hello())
}
