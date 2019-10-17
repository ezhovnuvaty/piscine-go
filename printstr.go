package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for a, char := range str {
		z01.PrintRune(char)
		a++
	}
}
