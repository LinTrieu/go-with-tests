// Dependency Injection
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection
// There is an article by Martin Fowler: https://martinfowler.com/articles/injection.html
package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}
