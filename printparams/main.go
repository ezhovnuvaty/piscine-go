package main

import (
	"github.com/01-edu/z01"
	"os"
)
func main()  {
	args := os.Args[1:]
	for _, word := range args {
		for _, letter := range word {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}
