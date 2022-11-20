// Mocking
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprint(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}

// define the dependency as an interface, use real Sleeper in main and a Spy in tests
// by using an interface, Countdown is obvlivious and adds flexibility for the caller
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (ds *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// TODO: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#still-some-problems
