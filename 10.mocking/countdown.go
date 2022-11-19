// Mocking
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
