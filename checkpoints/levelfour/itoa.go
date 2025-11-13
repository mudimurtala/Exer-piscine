// ## itoa
// ### Instructions
// - Write a function that simulates the behavior of the `Itoa` function in Go. `Itoa` transforms a number represented as an`int` in a number represented as a `string`.
// - For this exercise the handling of the signs + or - **does have** to be taken into account.
// ## Expected function
// ```go
// func Itoa(n int) string {
// }
// ```
// ### Usage
// Here is a possible program to test your function :
// ```go
// package main
// import (
// 	"fmt"
// 	"piscine"
// )
// func main() {
//     fmt.Println(piscine.Itoa(12345))
//     fmt.Println(piscine.Itoa(0))
//     fmt.Println(piscine.Itoa(-1234))
//     fmt.Println(piscine.Itoa(987654321))
// }
// ```
// And its output :
// ```console
// $ go run .
// 12345
// 0
// -1234
// 987654321
// $
// ```
// ### Notions
// - [strconv/Itoa](https://pkg.go.dev/strconv#Itoa)



package main
import (
	"fmt"
)

func Itoa(n int) string {
	// 1️⃣ Handle zero
	if n == 0 {
		return "0"
	}

	isNegative := false

	// 2️⃣ Handle negative numbers
	if n < 0 {
		isNegative = true
		n = -n
	}

	result := ""

	// 3️⃣ Extract digits from right to left
	for n > 0 {
		digit := n % 10                   // get last digit
		char := '0' + rune(digit)         // convert to rune
		result += string(char)            // add to result
		n = n / 10                        // remove last digit
	}

	// 4️⃣ Reverse the result (because digits were added in reverse order)
	runes := []rune(result)
	length := len(runes)
	i := 0               // Start from the beginning
	j := length - 1      // Start from the end
	for i < j {
		temp := runes[i]
		runes[i] = runes[j]
		runes[j] = temp
		i++
		j--
	}
	result = string(runes)

	// 5️⃣ Add minus sign if negative
	if isNegative {
		result = "-" + result
	}

	return result
}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}
