package main

import (
    "fmt"
)

func main() {
    // We print the result of the function call
    fmt.Println(removeSpace("Mudi , where are you going ?"))
}

func removeSpace(s string) string {
    runes := []rune(s)
    // 1. THE BUCKET: We create an empty slice to hold the characters we want to keep
    var result []rune

    for i := 0; i < len(runes); i++ {
        // 2. THE FILTER: Check if current is space AND next is punctuation
        if runes[i] == ' ' && i+1 < len(runes) && isPunctuation(runes[i+1]) {
            // We do nothing here (this is the "skip")
            continue 
        }

        // 3. THE COLLECTION: If we didn't skip, we add the character to our bucket
        result = append(result, runes[i])
    }

    // 4. THE CONVERSION: Turn the bucket back into a string
    return string(result)
}

// 5. THE HELPER: A dedicated function to identify our magnets
func isPunctuation(r rune) bool {
    // Switch is cleaner than multiple 'if' statements
    switch r {
    case ',', '.', '!', '?', ':', ';':
        return true
    }
    return false
}