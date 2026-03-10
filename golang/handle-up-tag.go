package main

import (
	"fmt"
	"strings"
)

func main() {
	// Test case 1: Single tag
	test1 := []string{"(up)", "hi", "mudi", "(up)", "how", "are", "you"}
	fmt.Println("Before:", test1)
	
	result1 := handleUpTag(test1)
	fmt.Println("After: ", result1)

	// Test case 2: Consecutive tags (This tests the 'i--' logic)
	test2 := []string{"hello", "world", "(up)", "(up)"}
	fmt.Println("\nBefore:", test2)
	fmt.Println("After: ", handleUpTag(test2))
}


func handleUpTag(words []string) []string {
    for i := 0; i < len(words); i++ {
        if words[i] == "(up)" {
            if i > 0 {
                words[i-1] = strings.ToUpper(words[i-1])
            }
            words = append(words[:i], words[i+1:]...)
            i-- 
        }
    }
    return words
}