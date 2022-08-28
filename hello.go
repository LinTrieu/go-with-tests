package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

// outside world
func main() {
	fmt.Println(Hello("Chris"))
}

// domain code
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}
