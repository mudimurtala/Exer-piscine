package main

import "fmt"

func main() {
	fmt.Println(changeFirstLetterToX("m?ello"))
	

}

func changeFirstLetterToX (s string) string {
	runes := []rune(s)
	fmt.Printf("%c\n", runes[1])
	if len(runes) > 0 {
		runes[0] = 'y'
	}
	return string(runes)
}
