package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	for index := range os.Args {
		count = index
	}
	for i := count; i > 0; i-- {
		for _, l2 := range os.Args[i] {
			z01.PrintRune(l2)
		}
		z01.PrintRune(10)
	}
}
