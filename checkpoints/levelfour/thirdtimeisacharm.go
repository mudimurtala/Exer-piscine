// ## thirdtimeisacharm

// ### Instructions

// Write a function `ThirdTimeIsACharm()` that takes a `string` as an argument and returns another `string` with every third character.

// - Return the output followed by a newline `\n`.
// - If the `string` is empty, return a newline `\n`.
// - If there is no third character, return a newline `\n`.

// ### Expected function

// ```go
// func ThirdTimeIsACharm(str string) string {

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
// 	fmt.Print(ThirdTimeIsACharm("123456789"))
// 	fmt.Print(ThirdTimeIsACharm(""))
// 	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
// 	fmt.Print(ThirdTimeIsACharm("12"))
// }
// ```

// And its output:

// ```console
// $ go run . | cat -e
// 369$
// $
// b e$
// $
// ```




package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	if str == "" {
		return "\n"
	}

	result := ""
	for i := 0; i < len(str); i++ {
		if (i + 1)%3 == 0 {
			result += string(str[i])
		}
	}

	if result == "" {
		return "\n"
	}

	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
