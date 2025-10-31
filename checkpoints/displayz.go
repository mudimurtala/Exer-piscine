package main

import "os"

// ✅ Imports the os package
// ✅ Starts at the main() function
// ✅ Prints the letter z, followed by a new line, directly to the terminal

func main() {
	os.Stdout.WriteString("z" + "\n")
}
