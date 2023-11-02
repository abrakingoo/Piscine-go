package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, char := range args[1:] {
		for i := 0; i < len(char); i++ {
			z01.PrintRune(rune(char[i]))
		}
		z01.PrintRune('\n')
	}
}
