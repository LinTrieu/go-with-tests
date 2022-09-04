package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 4, 5}
	got := Sum(numbers)
	expected := 10

	if expected != got {
		t.Errorf("Expected: %v, Got: %v, Given: %v", expected, got, numbers)
	}
}
