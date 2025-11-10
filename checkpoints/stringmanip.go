package main

import "fmt"

func StringManip(str string) string {
	runes := []rune(str)

	for i := 0; i < len(str); i++ {
		if i == 0 || runes[i - 1] == ' ' {
			if runes[i] != ' ' {
				runes[i] = 'R'
			}
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(StringManip("How are you"))
	fmt.Println(StringManip("What is your name"))
	
}
