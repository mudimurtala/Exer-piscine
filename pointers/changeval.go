package main

import "fmt"

func square(num *int) {
	*num = (*num) * (*num)
}

func main() {
	x := 9
	square(&x)
	fmt.Println(x)
}