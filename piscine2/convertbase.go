package main

import "fmt"

func NumFromBase(num int, base string) string {
	if num == 0 {
		return string(base[0])
	}
	mod := len(base)
	var result string
	for num > 0 {
		result = string(base[num%mod]) + result
		num /= mod
	}
	return result
}

func Index(s string, c string) int {
	for i, r := range s {
		if string(r) == c {
			return i
		}
	}
	return -1
}

func AtoiBase(str string, base string) int {
	if len(base) < 2 || Index(base, "+") >= 0 || Index(base, "-") >= 0 {
		return 0
	}
	for i1, char1 := range base {
		for i2, char2 := range base {
			if i1 != i2 && char1 == char2 {
				return 0
			}
		}
	}
	baseLen := len(base)
	result := 0

	for _, c := range str {
		result = result*baseLen + Index(base, string(c))
	}
	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AtoiBase(nbr, baseFrom)
	return NumFromBase(n, baseTo)
}

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result) // ✅ 43
}
