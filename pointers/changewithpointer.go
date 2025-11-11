package main

import "fmt"

func main() {
	str := "Go is fun!"
	p := &str

	*p = "Go is awesome!"
	fmt.Println(str)
}
