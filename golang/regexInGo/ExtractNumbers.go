// ### 🟢 Exercise 3 — Extract All Numbers
// Write a function `ExtractNumbers(s string) []string` that finds and returns all sequences of digits in a string.
// ```
// ExtractNumbers("I have 3 cats and 12 dogs")  → ["3", "12"]
// ExtractNumbers("no numbers here")            → []
// ExtractNumbers("room 101 on floor 3")        → ["101", "3"]
// ```

package main

import (
	"fmt"
	"regexp"
)

func ExtractNumbers(s string) []string {
	re := regexp.MustCompile(`\d+`)
	return re.FindAllString(s, -1)
}

func main() {
	fmt.Printf("%q\n", ExtractNumbers("I have 3 cats and 12 dogs"))
	fmt.Printf("%q\n", ExtractNumbers("no numbers here"))
	fmt.Printf("%q\n", ExtractNumbers("room 101 on floor 3"))
}
