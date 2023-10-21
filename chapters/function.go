package main

import "fmt"

func main() {
	fmt.Printf(multipleResults("Go", "Quan"))
}

// return x, y (used only in short functions)
//`naked` return x, y (used only in short functions)
func multipleResults(a, b string) (x, y string) {
	x = a + "Hello"
	y = b + "World"

	return
}
