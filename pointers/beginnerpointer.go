package main

import "fmt"

func main() {
	z := 25
	y := &z

	fmt.Println("z:", z)
	fmt.Println("Address of z:", y)
	fmt.Println("Value stored in address:", *y)
}
