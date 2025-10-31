package main

import "fmt"

// ✅ Final Summary (In Simple Words)
// This program:
// ✔ Defines a function IsAlpha
// ✔ Checks each character of a string
// ✔ If any character is not a letter or number → returns false
// ✔ If all characters are valid → returns true
// ✔ Tests the function with different strings and prints the results

func IsAlpha(s string) bool {
	array := []rune(s)

	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(array); i++ {
		if (array[i] >= 'A' && array[i] <= 'Z') || (array[i] >= 'a' && array[i] <= 'z') || (array[i] >= '0' && array[i] <= '9') {
			continue
		} else {
			return false
		}
	}

	return true
}


func main() {
    fmt.Println(IsAlpha("Hello! How are you?"))
    fmt.Println(IsAlpha("HelloHowareyou"))
    fmt.Println(IsAlpha("What's this 4?"))
    fmt.Println(IsAlpha("Whatsthis4"))
    fmt.Println(IsAlpha(""))
	fmt.Println(len("😊"))
	fmt.Println(len([]rune("😊")))
}