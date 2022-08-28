package main

import (
	"fmt"
)

// outside world
func main() {
	fmt.Println(Hello())
}

// domain code
func Hello() string {
	return "Hello, World"
}
