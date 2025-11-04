// ## count-character
// ### Instructions
// write a function that takes a string and a character as 
// arguments and returns the number of times the character appears in the string.
// - if the character is not in the string return 0
// - if the string is empty return 0
// ### Expected Function
// ```go
// func CountChar(str string, c rune) int {
//     ...
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
// 	fmt.Println(CountChar("Hello World", 'l'))
// 	fmt.Println(CountChar("5  balloons", 5))
// 	fmt.Println(CountChar("   ", ' '))
// 	fmt.Println(CountChar("The 7 deadly sins", '7'))
// }
// ```
// And its output :
// ```console
// $ go run .
// 3
// 0
// 3
// 1
// ```



package main

import (
	"fmt"
)

func CountChar(str string, c rune) int {
	count := 0

	if str == "" {
		return 0
	}

	for i := 0; i < len(str); i++ {
		char := rune(str[i])
		if char == c {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
	fmt.Println(CountChar("", '7'))
}
