// Dependency Injection
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection
package main

import (
	"bytes"
	"testing"
)

// we need inject the dependency of printing.
// Our function doesn't care where or how the printing happens
// so we should accept an interface rather than a concrete type.

// The Buffer type implements the Writer interface, because it has the method Write(p []byte) (n int, err error).
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Evelyn")

	got := buffer.String()
	want := "Hello, Evelyn"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
