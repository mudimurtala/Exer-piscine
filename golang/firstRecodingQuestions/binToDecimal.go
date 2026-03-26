package main

import (
	"fmt"
	"strconv"
)


func binToDecimal(binStr string) (int64, error) {
	var result int64 = 0
	var err error

	for _, ch := range binStr {
		// Convert character to string, then to integer
		var digit int64
		digit, err = strconv.ParseInt(string(ch), 10, 64)
		if err != nil {
			return 0, err
		}

		// Check if it's valid binary (only 0 or 1)
		if digit != 0 && digit != 1 {
			return 0, fmt.Errorf("Invalid binary digit: %c", ch)
		}

		// Core logic
		result = result*2 + digit
	}

	return result, nil
}


func main() {
	res, err := binToDecimal("10")
	fmt.Printf("binToDecimal(\"10\") -> %v, %v\n", res, err)

	res, err = binToDecimal("1010")
	fmt.Printf("binToDecimal(\"1010\") -> %v, %v\n", res, err)

	res, err = binToDecimal("11111111")
	fmt.Printf("binToDecimal(\"11111111\") -> %v, %v\n", res, err)
}
