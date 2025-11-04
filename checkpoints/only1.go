package main

import "github.com/01-edu/z01"

// ## only1
// ### Instructions
// Write a program that displays a `1` character on the standard output. (and nothing else)
// $ go run .
// 1
// $ go run . "a" "b"
// 1
// $ go run . "a" "b" "c"
// 1

func main() {
	z01.PrintRune('1')
	z01.PrintRune(10)
}
