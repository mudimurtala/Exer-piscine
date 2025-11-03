package main

import "os"

// ## displaylastparam
// ### Instructions
// Write a program that displays its last argument, if there is one.
// ### Usage
// $ go run . hello there
// there
// $ go run . "hello there" how are you
// you
// $ go run . "hello there"
// hello there
// $ go run .
// $


// ✅ Summary in plain English
// This program:
// ✔ Checks if an argument is provided
// ✔ Finds the last argument typed
// ✔ Prints that last argument on a new line
// ✔ Does nothing if no argument is provided

func main() {
	if len(os.Args) < 2 {
		return
	}
	os.Stdout.WriteString(os.Args[len(os.Args) - 1] + "\n")
}
