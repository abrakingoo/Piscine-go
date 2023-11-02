package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for i := 0; i < len(arguments)-1; i++ {
		for j := 0; j < len(arguments)-i-1; j++ {
			if compare(arguments[j], arguments[j+1]) > 0 {
				arguments[j], arguments[j+1] = arguments[j+1], arguments[j]
			}
		}
	}

	for _, char := range arguments {
		z01.PrintRune(rune(char[0]))
		z01.PrintRune('\n')
	}

}

func compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
