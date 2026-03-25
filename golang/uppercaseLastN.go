Q1 - Write a function that fixes spacing around punctuation for a simple case 
Given a slice of tokens (words + punctuation), remove the space before punctuation marks. 

// Requirement: Please use a variable named 'result' for the output string, 
// and loop with 'i, token'.

// package main

// import (
// 	"fmt"
// )

// // Example - Input: ["hello", ",", "world", "!"] -> Expected Output: "hello, world!"
// func joinWithPunctuation(tokens []string) string {
  
// }

// func main() {
//   tokens := []string{"hello", ",", "world", "!"}
//   fmt.Printf("joinWithPunctuation(...) -> %q\n", joinWithPunctuation(tokens))
// }

// func isPunctuation(s string) bool {
//  punctuations := []string{".", ",", "!", "?", ":", ";"}
//  for _, p := range punctuations {
//  if s == p {
//  return true
//  }
//  }
//  return false
// }








Q2 - Write a function that converts a binary string to decimal
Convert a binary string to its decimal integer value.

// Requirement:
// You must use the variable names 'result' and 'err' inside your function.

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// Example 1 - Input: "10" -> Expected Output: 2
// Example 2 - Input: "1010" -> Expected Output: 10
// Example 3 - Input: "11111111" -> Expected Output: 255
// func binToDecimal(binStr string) (int64, error) {

// }

// func main() {
//   res, err := binToDecimal("10")
//   fmt.Printf("binToDecimal(\"10\") -> %v, %v\n", res, err)
//   res, err = binToDecimal("11111111")
//   fmt.Printf("binToDecimal(\"11111111\") -> %v, %v\n", res, err)
// }




Q3 - Write a function that converts a hexadecimal string to its decimal integer value
Without using strconv.ParseInt directly ; use it, but explain what base and bitSize mean.

// Requirement:
// You must use the variable names 'result' and 'err' inside your function.

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// // Example 1 - Input: "1E"  -> Expected Output: 30
// // Example 2 - Input: "FF"  -> Expected Output: 255
// func hexToDecimal(hexStr string) (int64, error) {

// }

// func main() {
//   res, err := hexToDecimal("1E")
//   fmt.Printf("hexToDecimal(\"1E\") -> %v, %v\n", res, err)
//   res, err = hexToDecimal("FF")
//   fmt.Printf("hexToDecimal(\"FF\") -> %v, %v\n", res, err)
// }




Q4 - Explain strings.Fields()
// Explain what strings.Fields() does and why it's useful in this project 




Q5 - Write a function that checks if a string is a punctuation mark from the projects list
Check if a string is a punctuation mark from the projects list.

// Requirement:
// Please use a variable named 'punctuations' for the array, and 'p' in the loop.

// package main

// import (
// 	"fmt"
// )

// // Example 1 - Input: "," -> Expected Output: true
// // Example 2 - Input: "!" -> Expected Output: true
// // Example 3 - Input: "x" -> Expected Output: false
// func isPunctuation(s string) bool {

// }

// func main() {
//   fmt.Printf("isPunctuation(\",\") -> %v\n", isPunctuation(","))
//   fmt.Printf("isPunctuation(\"x\") -> %v\n", isPunctuation("x"))
// }






Q6 - Write a function that fixes spacing inside single quotes
Fix spacing inside single quotes.

// Requirement:
// Declare the exact variables: 'result', 'match', and 'inner'.

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// // Example 1 - Input: "' awesome '" -> Expected Output: "'awesome'"
// // Example 2 - Input: "' hello world '" -> Expected Output: "'hello world'"
// func fixSingleQuotes(text string) string {

// }

// func main() {
//   fmt.Printf("fixSingleQuotes(...) -> %q\n", fixSingleQuotes("' awesome '"))
//   fmt.Printf("fixSingleQuotes(...) -> %q\n", fixSingleQuotes("' hello world '"))
// }


Q7 - Write a function that takes a slice of words and applies uppercase to the last N words
Take a slice of words and apply uppercase to the last N words.

// Requirement:
// Please use a variable named 'start' to track the starting index.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // Example - Input: words = ["this", "is", "so", "exciting"], n = 2
// // Expected Output: ["this", "is", "SO", "EXCITING"]
// func uppercaseLastN(words []string, n int) []string {

// }

// func main() {
//   words := []string{"this", "is", "so", "exciting"}
//   fmt.Printf("uppercaseLastN(..., 2) -> %q\n", uppercaseLastN(words, 2))
// }




Q8 - Write a function that determines whether to use "a" or "an" before a given word
Determine whether to use "a" or "an" before a given word.

// Requirement:
// Name variables exactly: 'firstLetter', 'vowelsAndH' and loop using 'v'.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // Example 1 - Input: "apple" -> Expected Output: "an"
// // Example 2 - Input: "horse" -> Expected Output: "an"
// // Example 3 - Input: "book" -> Expected Output: "a"
// // Example 4 - Input: "honest" -> Expected Output: "an" (starts with silent h)
// func aOrAn(nextWord string) string {

// }

// func main() {
//   fmt.Printf("aOrAn(\"apple\") -> %q\n", aOrAn("apple"))
//   fmt.Printf("aOrAn(\"horse\") -> %q\n", aOrAn("horse"))
//   fmt.Printf("aOrAn(\"honest\") -> %q\n", aOrAn("honest"))
// }