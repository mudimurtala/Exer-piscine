package main

import "os"

// ✅ What does the program do overall?
// This program:
// Checks if the user typed something after the program name
// If yes → it prints what they typed
// If not → it does nothing and exits quietly

func main() {
	if len(os.Args) < 2 {
		return
	}
	os.Stdout.WriteString(os.Args[1] + "\n")
}
