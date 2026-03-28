// ### 🟢 Exercise 2 — Starts With a Capital Letter?
// Write a function `StartsWithCapital(s string) bool` that returns `true` if the string begins with an uppercase letter (A–Z).
// ```
// StartsWithCapital("Hello")   → true
// StartsWithCapital("hello")   → false
// StartsWithCapital("123abc")  → false
// ```

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
