// Maps
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps
package main

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}
