package main

import "fmt"

func main() {
	fmt.Printf(nakedResults("Go 1.17", "Quan"))
}

// return x, y (used only in short functions)
// `naked` return x, y (used only in short functions)
func nakedResults(a, b string) (x, y string) {
	x = a + "Hello"
	y = b + "World"

	return
}
