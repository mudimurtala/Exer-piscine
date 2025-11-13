// ## printmemory

// ### Instructions

// Write a function that takes `(arr [10]byte)`, and displays the memory as in the example.

// After displaying the memory the function must display all the ASCII graphic characters. 
// The non printable characters must be replaced by a dot.

// The ASCII graphic characters are any characters intended to be written, printed, 
// or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.

// ### Expected function

// ```go
// func PrintMemory(arr [10]byte) {

// }
// ```

// ### Usage

// Here is a possible program to test your function :

// ```go
// package main

// import "piscine"

// func main() {
// 	piscine.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
// }
// ```

// And its output :

// ```console
// $ go run . | cat -e
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
// $
// ```



package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	count := 0 // to track how many bytes printed per line

	for i := 0; i < len(arr); i++ {
		printHex(arr[i])

		count++
		// Only print a space if this isn't the 4th element or the last element
		if count < 4 && i != len(arr)-1 {
			z01.PrintRune(' ')
		}

		// Start new line every 4 bytes or at the end
		if count == 4 || i == len(arr)-1 {
			z01.PrintRune('\n')
			count = 0
		}
	}

	// Print ASCII characters (printable range: 32–126)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

// Helper to print a byte as hexadecimal (two digits)
func printHex(b byte) {
	hexChars := "0123456789abcdef"
	high := b / 16
	low := b % 16
	z01.PrintRune(rune(hexChars[high]))
	z01.PrintRune(rune(hexChars[low]))
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
