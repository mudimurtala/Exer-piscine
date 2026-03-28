// ### 🟢 Exercise 1 — Does It Contain a Number?
// Write a function `ContainsNumber(s string) bool` that returns `true` if the input string contains at least one digit, and `false` otherwise.
// ```
// ContainsNumber("hello123")  → true
// ContainsNumber("hello")     → false
// ContainsNumber("abc9xyz")   → true
// ```



/* Non Regex Solution */
// package main
// import ("fmt")
// func ContainsNumber(s string) bool {
// 	for i := 0; i < len(s); i++ {
// 		if s[i] >= '0' && s[i] <= '9' {
// 			return true
// 		}
// 	}
// 	return false
// }
// func main() {
// 	fmt.Println(ContainsNumber("hello123"))
// 	fmt.Println(ContainsNumber("hello"))
// 	fmt.Println(ContainsNumber("abc9xyz"))
// }





/* Regex Solution */
package main

import (
	"fmt"
	"regexp"
)

func ContainsNumber(s string) bool {
	re := regexp.MustCompile(`\d+`)
	return re.MatchString(s)
}

func main() {
	fmt.Println(ContainsNumber("hello123"))
	fmt.Println(ContainsNumber("hello"))
	fmt.Println(ContainsNumber("abc9xyz"))
}
