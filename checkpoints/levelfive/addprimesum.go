// ## addprimesum

// ### Instructions

// Write a program that takes a positive integer as argument and displays 
// the sum of all prime numbers inferior or equal to it followed by a newline (`'\n'`).
// - If the number of arguments is different from 1, or if the argument is 
// not a positive number, the program displays `0` followed by a newline.

// ### Usage

// ```console
// $ go run . 5
// 10
// $ go run . 7
// 17
// $ go run . -2
// 0
// $ go run . 0
// 0
// $ go run .
// 0
// $ go run . 5 7
// 0
// $
// ```




package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	// Step 1: Check number of arguments
	if len(os.Args) != 2 {
		printNumber(0)
		return
	}

	// Step 2: Convert argument to integer
	n, ok := atoi(os.Args[1])
	if !ok || n <= 0 {
		printNumber(0)
		return
	}

	// Step 3: Loop from 2 to n and sum primes
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	// Step 4: Print the sum
	printNumber(sum)
}

// --------------------------------------------------------
// Helper: Check if a number is prime (simple approach)
// --------------------------------------------------------
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ { // check up to √n
		if n%i == 0 {
			return false
		}
	}
	return true
}

// --------------------------------------------------------
// Helper: Convert string to int (BEGINNER VERSION)
// --------------------------------------------------------
func atoi(s string) (int, bool) {
	sign := 1
	result := 0

	// Handle optional + or -
	index := 0
	if len(s) > 0 {
		if s[0] == '-' {
			sign = -1
			index++
		} else if s[0] == '+' {
			index++
		}
	}

	// Loop through characters
	for ; index < len(s); index++ {
		ch := s[index]
		if ch < '0' || ch > '9' {
			return 0, false
		}
		result = result*10 + int(ch-'0')
	}

	return result * sign, true
}

// --------------------------------------------------------
// Helper: Print number using z01.PrintRune
// --------------------------------------------------------
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	digits := []rune{}

	for n > 0 {
		d := n % 10
		digits = append(digits, rune(d+'0'))
		n /= 10
	}

	// reverse digits
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}

	z01.PrintRune('\n')
}
