package main

import "github.com/01-edu/z01"

// ## onlyb
// ### Instructions
// Write a program that displays a `B` character on the standard output. (and nothing else)
// $ go run .
// B
// $ go run . "a" "b"
// B
// $ go run . "a" "b" "c"
// B

func main() {
	z01.PrintRune('B')
	z01.PrintRune(10)
}
