package main

import (
	"fmt"
	"strconv"
)


func hexToDecimal(hexStr string) (int64, error) {
	var result int64
	var err error

	// Convert hex string to decimal
	result, err = strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	res, err := hexToDecimal("1E")
	fmt.Printf("hexToDecimal(\"1E\") -> %v, %v\n", res, err)

	res, err = hexToDecimal("FF")
	fmt.Printf("hexToDecimal(\"FF\") -> %v, %v\n", res, err)
}