package main

import "fmt"

func main() {
	s := []int{1, 2}
	t := 3

	s = append(s, t)
	fmt.Println(s)
}