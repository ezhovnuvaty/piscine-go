package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args[0]
	for _, str := range name {
		z01.PrintRune(str)
	}
	z01.PrintRune(10)
}
