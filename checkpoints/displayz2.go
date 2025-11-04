// ## displayz
// ### Instructions
// Write a program that takes a `string`, and displays the first `z` character it encounters in it, followed by a newline (`'\n'`).
// If there are no `z` characters in the string, the program just writes `z` followed by a newline (`'\n'`).
// If the number of arguments is not 1, the program displays a `z` followed by a newline (`'\n'`).
// ### Usage
// ```console
// $ go run . "xyz"
// z
// $ go run . "bcvbvZ"
// z
// $ go run . "nbv"
// z
// $
// ```

package main

import "os"

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("z" + "\n")
		return
	}

	input := os.Args[1]
	found := false

	for _, char := range input {
		if char == 'z' {
			os.Stdout.WriteString("z" + "\n")
			found = true
		}
	}

	if !found {
		os.Stdout.WriteString("z" + "\n")
	}
}
