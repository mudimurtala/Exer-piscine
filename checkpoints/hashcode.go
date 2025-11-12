// ## hashcode
// ### Instructions
// Write a function called `HashCode()` that takes a `string` as an argument
//  and returns a new **hashed** `string`.
// - The hash equation is computed as follows:
// `(ASCII of current character + size of the string) % 127, ensuring the result falls
//  within the ASCII range of 0 to 127.`
// - If the resulting character is unprintable add `33` to it.
// ### Expected function
// ```go
// func HashCode(dec string) string {
// }
// ```
// ### Usage
// Here is a possible program to test your function:
// ```go
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }
// ```
// And its output:
// ```console
// $ go run .
// B
// CD
// EDF
// Spwwz+bz}wo
// ```



package main

import (
	"fmt"
)

func HashCode(dec string) string {
	result := ""
	for i := 0; i < len(dec); i++ {
		char := dec[i]
		hv := (int(char) + len(dec)) % 127
		if hv < 32 {
			hv = hv + 33
		}
		result = result + string(rune(hv))
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
	fmt.Println(HashCode("MUDI"))
	fmt.Println(HashCode("mudi"))
	fmt.Println(HashCode("my name is mudi Muriata"))
}