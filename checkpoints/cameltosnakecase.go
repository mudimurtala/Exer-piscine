// ## cameltosnakecase
// ### Instructions
// Write a function that converts a `string` from `camelCase` to `snake_case`.
// - If the `string` is empty, return an empty `string`.
// - If the `string` is not `camelCase`, return the `string` unchanged.
// - If the `string` is `camelCase`, return the `snake_case` version of the `string`.
// For this exercise you need to know that `camelCase` has two different writing alternatives that will be accepted:
// - lowerCamelCase
// - UpperCamelCase
// Rules for writing in `camelCase`:
// - The word does not end on a capitalized letter (CamelCasE).
// - No two capitalized letters shall follow directly each other (CamelCAse).
// - Numbers or punctuation are not allowed in the word anywhere (camelCase1).
// ### Expected function
// ```go
// func CamelToSnakeCase(s string) string{
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
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }
// ```
// // And its output:
// ```console
// $ go run .
// Hello_World
// hello_World
// camel_Case
// CAMELtoSnackCASE
// camel_To_Snake_Case
// hey2
// ```



package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	array := []rune(s)

	for i := 0; i < len(array); i++ {
		r := array[i]
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			return s
		}
	}

	last := array[len(array) - 1]
	if last >= 'A' && last <= 'Z' {
		return s
	}

	prevUpper := false
	for i := 0; i < len(array); i++ {
		r := array[i]
		isUpper := (r >= 'A' && r <= 'Z')
		if isUpper && prevUpper {
			return s
		}
		prevUpper = isUpper
	}

	out := make([]rune, 0, len(array) * 2)
	for i := 0; i < len(array); i++ {
		r := array[i]
		if r >= 'A' && r <= 'Z' && i != 0 {
			out = append(out, '_')
		}
		out = append(out, r)
	}

	return string(out)
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}