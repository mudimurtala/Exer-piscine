//Q3 - Write a function that converts a binary string to decimal
//Convert a binary string to its decimal integer value.
//Requirement:
//You must use the variable names 'result' and 'err' inside your function.

package main

import (
	"fmt"
	"strconv"
)

//Example 1 - Input: "10" -> Expected Output: 2
//Example 2 - Input: "1010" -> Expected Output: 10
//Example 3 - Input: "11111111" -> Expected Output: 255

func binToDecimal(binStr string) (int64, error) {

}

func main() {
  res, err := binToDecimal("10")
  fmt.Printf("binToDecimal(\"10\") -> %v, %v\n", res, err)
  res, err = binToDecimal("11111111")
  fmt.Printf("binToDecimal(\"11111111\") -> %v, %v\n", res, err)
}