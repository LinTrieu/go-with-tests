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
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprint(out, i)
		fmt.Fprint(out, "\n")
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, "Go!")
}
