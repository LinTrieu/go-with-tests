// Mocking
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

package main

import (
	"bytes"
	"testing"
)

// Our software needs to print to stdout
// The bytes.Buffer type implements the Writer interface
// So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Countrdown
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
	2
	1
	Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
