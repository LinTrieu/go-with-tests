// Structs, methods & interfaces
// https: //quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces
package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
