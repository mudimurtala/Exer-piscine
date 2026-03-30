// ### 🟢 Exercise 2 — Starts With a Capital Letter?
// Write a function `StartsWithCapital(s string) bool` that returns `true` if the string begins with an uppercase letter (A–Z).
// ```
// StartsWithCapital("Hello")   → true
// StartsWithCapital("hello")   → false
// StartsWithCapital("123abc")  → false
// ```



/* Non Regex Solution*/
// package main
// import (
// 	"fmt"
// )
// func StartsWithCapital(s string) bool {
// 	if s[0] >= 'A' && s[0] <= 'Z' {
// 		return true
// 	}
// 	return false
// }
// func main() {
// 	fmt.Println(StartsWithCapital("Hello"))
// 	fmt.Println(StartsWithCapital("hello"))
// 	fmt.Println(StartsWithCapital("123abc"))
// }



/* Regex Solution*/
package main

import (
	"fmt"
	"regexp"
)

func StartsWithCapital(s string) bool {
	re := regexp.MustCompile(`^[A-Z]`)
	return re.MatchString(s)
}

func main() {
	fmt.Println(StartsWithCapital("Hello"))
	fmt.Println(StartsWithCapital("hello"))
	fmt.Println(StartsWithCapital("123abc"))
}