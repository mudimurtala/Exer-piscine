package main

import "os"

func main() {
	if len(os.Args) < 2 {
		return
	}
	os.Stdout.WriteString(os.Args[len(os.Args) - 1] + "\n")
}
