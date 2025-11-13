// ## findprevprime
// ### Instructions
// Write a function that returns the first prime number that is 
// equal or inferior to the `int` passed as parameter.
// If there are no primes inferior to the `int` passed as parameter
// the function should return 0.
// ### Expected function
// ```go
// func FindPrevPrime(nb int) int {// }
// ```
// ### Usage
// Here is a possible program to test your function :
// ```go
// package main
// import "fmt"
// func main() {
// 	fmt.Println(FindPrevPrime(5))
// 	fmt.Println(FindPrevPrime(4))
// }
// ```
// And its output :
// ```console
// $ go run .
// 5
// 3
// $
// ```



package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 { // no prime less than 2
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
	fmt.Println(FindPrevPrime(10))
	fmt.Println(FindPrevPrime(1))
}
