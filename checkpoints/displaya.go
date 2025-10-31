package main

import "os"

// ✅ Imports the os package
// ✅ Starts at the main() function
// ✅ Prints the letter a, followed by a new line, directly to the terminal

func main() {
	os.Stdout.WriteString("a" + "\n")
}
