package main

import "fmt"

// ✅ What does the program do overall?
// ✔ Loops through each character
// ✔ Verifies if it is between 'a' and 'z'
// ✔ If anything is not lowercase → stop and return false
// ✔ Otherwise → true
// ➡️ It is a lowercase-only string checker

func IsLower(s string) bool {
    array := []rune(s)

    for i := 0; i < len(array); i++ {
        if array[i] < 'a' || array[i] > 'z' {
            return false
        }
    }

    return true
}

func main() {
    fmt.Println(IsLower("hello"))
    fmt.Println(IsLower("hello!"))
}
