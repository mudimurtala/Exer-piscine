// ## checknumber
// ### Instructions
// Write a function that takes a `string` as an argument and returns `true` if the string 
// contains any number, otherwise return `false`.
// ### Expected function
// ```go
// func CheckNumber(arg string) bool {
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
// 	fmt.Println(CheckNumber("Hello"))
// 	fmt.Println(CheckNumber("Hello1"))
// }
// ```
// And its output:
// ```console
// $ go run .
// false
// true
// $
// ```


package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
	for _, num := range arg {
		if num >= '0' && num <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}