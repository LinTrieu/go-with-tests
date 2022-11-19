// Dependency Injection
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection
// There is an article by Martin Fowler: https://martinfowler.com/articles/injection.html
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

// fmt.Fprintf allows you to pass in an io.Writer which both os.Stdout and bytes.Buffer implement
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "worlds!")
}
