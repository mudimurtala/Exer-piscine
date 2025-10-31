package main

import "github.com/01-edu/z01"

func PrintString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func IsNumeric(s string) {
	if len(s) == 0 {
		PrintString("false\n")
		return
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			PrintString("false\n")
			return
		}
	}

	PrintString("true\n")
}

func main() {
	IsNumeric("010203")
	IsNumeric("01,02,03")
	IsNumeric("")
}
