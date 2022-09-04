package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [2]int{1, 4}
	got := sum(numbers)
	expected := 5

	if expected != got {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}
