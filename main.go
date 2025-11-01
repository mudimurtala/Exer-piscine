package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printstr("NV")
		return
	}
	baseRunes := []rune(base)
	bSize := int64(len(baseRunes))
	n := int64(nbr)
	if n == 0 {
		z01.PrintRune(baseRunes[0])
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	var result []rune
	for n != 0 {
		mod := n % bSize
		if mod < 0 {
			mod = -mod
		}
		result = append([]rune{baseRunes[mod]}, result...)
		n /= bSize
	}
	for _, r := range result {
		z01.PrintRune(r)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if r == rune(base[j]) {
				return false
			}
		}
	}
	return true
}

func printstr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
