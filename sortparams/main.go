package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for i := 0; i < len(arguments); i++ {
		for j := i + 1; j < len(arguments); j++ {
			if arguments[i] > arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}

	for _, args := range arguments {
		for _, char := range args {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
