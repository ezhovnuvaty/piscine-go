package piscine

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	result := []rune(s)

	for index, value := range result {
		if index == len(result)-1 {
			return value
		}
	}
	return 0
}
