package main

import (
	"reflect"
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

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{3, 9, 1})
		want := []int{7, 10}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
