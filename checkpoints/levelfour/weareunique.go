// ## weareunique

// ### Instructions

// Write a function that takes two `strings`'s and returns the number 
// of characters that are not included in both, without repeating characters.

// - If there is no unique characters return `0`.
// - If both strings are empty return `-1`.

// ### Expected function

// ```go
// func WeAreUnique(str1 , str2 string) int {

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
// 	fmt.Println(WeAreUnique("foo", "boo"))
// 	fmt.Println(WeAreUnique("", ""))
// 	fmt.Println(WeAreUnique("abc", "def"))
// }
// ```

// And its output:

// ```console
// $ go run .
// 2
// -1
// 6
// ```




package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	all := str1 + str2 // combine both strings
	unique := ""       // to store already counted characters
	count := 0

	for _, ch := range all {
		c := string(ch)
        if contains(unique, c) {
			continue
		}

		in1 := contains(str1, c)
		in2 := contains(str2, c)

		if (in1 && !in2) || (!in1 && in2) {
			count++
			unique += c // mark it counted
		}
	}

	if count == 0 {
		return 0
	}
	return count
}

func contains(s, char string) bool {
	for _, ch := range s {
		if string(ch) == char {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
