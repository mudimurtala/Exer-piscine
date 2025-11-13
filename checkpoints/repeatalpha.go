// ## repeatalpha
// ### Instructions
// Write a function called `RepeatAlpha` that takes a `string` and displays it 
// repeating each alphabetical character as many times as its alphabetical index.
// `'a'` becomes `'a'`, `'b'` becomes `'bb'`, `'e'` becomes `'eeeee'`, etc...
// ### Expected Function
// ```go
// func RepeatAlpha(s string) string {
// }
// ```
// ### Usage
// Here is a possible program to test your function:
// ```go
// package main
// import (
// 	"fmt"
// 	"piscine"
// )
// func main() {
// 	fmt.Println(piscine.RepeatAlpha("abc"))
// 	fmt.Println(piscine.RepeatAlpha("Choumi."))
// 	fmt.Println(piscine.RepeatAlpha(""))
// 	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
// }
// ```
// And its output:
// ```console
// $ go run . | cat -e
// abbccc$
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// $
// abbacccaddddabba 01!$
// $
// ```



package main
import (
	"fmt"
)

func RepeatAlpha(s string) string {
	result := "" // this will store our final string

	for _, r := range s { // loop through each character
		repeatCount := 1 // by default, repeat once

		if r >= 'a' && r <= 'z' {
			repeatCount = int(r - 'a' + 1)
		} else if r >= 'A' && r <= 'Z' {
			repeatCount = int(r - 'A' + 1)
		}

		// Add the character "repeatCount" times to result
		for i := 0; i < repeatCount; i++ {
			result += string(r)
		}
	}

	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
