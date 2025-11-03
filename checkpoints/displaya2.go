// ## displaya
// ### Instructions
// Write a program that takes a `string`, and displays the first `a` character it 
// encounters in it, followed by a newline (`'\n'`).
// If there are no `a` characters in the string, the program just writes `a` 
// followed by a newline (`'\n'`).
// If the number of arguments is not 1, the program displays an `a` 
// followed by a newline (`'\n'`).
// ### Usage
// ```console
// $ go run . "abc"
// a
// $ go run . "bcvbvA"
// a
// $ go run . "nbv"
// a
// $
// ```


package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('a')
		z01.PrintRune(10)
		return
	}

	found := false

	for _, char := range os.Args[1] {
		if char == 'a' {
			z01.PrintRune('a')
			z01.PrintRune(10)
			found = true
			break
		}
	}

	if !found {
		z01.PrintRune('a')
		z01.PrintRune(10)
	}
}
