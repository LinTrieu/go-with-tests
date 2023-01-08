// Mocking
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

// define the dependency as an interface, use real Sleeper in main and a Spy in tests
// by using an interface, Countdown is obvlivious and adds flexibility for the caller
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// TODO: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#still-some-problems
