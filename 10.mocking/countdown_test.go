// Mocking
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

package main

import (
	"bytes"
	"testing"
)

// Spies are mocks to record how a dependency is used (record arguments, count calls, etc.)
// Sleep() keeps track of how many times it is called
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Our software needs to print to stdout
// The bytes.Buffer type implements the Writer interface
// So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Countrdown
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `321Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3, got %d", spySleeper.Calls)
	}
}
