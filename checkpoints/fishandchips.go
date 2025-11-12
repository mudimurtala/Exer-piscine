// ## fishandchips
// ### Instructions
// Write a function called `FishAndChips()` that takes an `int` and returns a `string`.
// - If the number is divisible by 2, print `fish`.
// - If the number is divisible by 3, print `chips`.
// - If the number is divisible by 2 and 3, print `fish and chips`.
// - If the number is negative return `error: number is negative`.
// - If the number is non divisible by 2 or 3 return `error: non divisible`.
// ### Expected function
// ```go
// func FishAndChips(n int) string {
// }
// ```
// ### Usage
// Here is a possible program to test your function:
// ```go
// package main
// import (
// 	"fmt"
// 	"piscine"
// )
// func main() {
// 	fmt.Println(piscine.FishAndChips(4))
// 	fmt.Println(piscine.FishAndChips(9))
// 	fmt.Println(piscine.FishAndChips(6))
// }
// ```
// And its output:
// ```console
// $ go run . | cat -e
// fish$
// chips$
// fish and chips$
// ```



package main
import (
	"fmt"
)

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if n % 2 == 0 && n % 3 == 0 {
		return "fish and chips"
	}

	if n % 2 == 0 {
		return "fish"
	}

	if n % 3 == 0 {
		return "chips"
	}

	return "error: non divisible"
}

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
	fmt.Println(FishAndChips(-8))
	fmt.Println(FishAndChips(5))
}
