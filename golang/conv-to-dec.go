package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Test Binary
	val, err := binToDec("10")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Binary 101 is:", val) // Should be 5
	}

	// Test Hex
	hexVal, err := hexToDec("1E")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hex 1E is:", hexVal) // Should be 30
	}

	// Test Invalid
	_, err = binToDec("102") 
	fmt.Println("Invalid Bin Test:", err) // Should show an error!
}

func binToDec(s string) (int64, error) {
	// We use 's' (the input), not a hardcoded string
	result, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, err // Return the error to the caller
	}
	return result, nil // Return the value and 'nil' (no error)
}

func hexToDec(s string) (int64, error) {
	result, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
