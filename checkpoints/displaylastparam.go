package main

import "os"

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
