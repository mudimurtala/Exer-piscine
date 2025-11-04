package main

import "github.com/01-edu/z01"

// ## onlyf
// ### Instructions
// Write a program that displays a `f` character on the standard output. (and nothing else)
// $ go run .
// f
// $ go run . "a" "b"
// f
// $ go run . "a" "b" "c"
// f

func main() {
	z01.PrintRune('f')
	z01.PrintRune(10)
}
