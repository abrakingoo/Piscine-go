package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	result := 0
	sign := 1

	for i, char := range s {
		if i == 0 && char == '-' {
			sign = -1
			continue
		} else if i == 0 && char == '+' {
			continue
		} else if char < '0' || char > '9' {
			return 0
		}
		result = result*10 + int(char-'0')
	}
	return result * sign
}

func findChar(char int) rune {
	count := 1
	if char >= 1 && char <= 26 {
		for i := 'a'; i <= 'z'; i++ {
			if count == char {
				return i
			}
			count++
		}
	}
	return ' '
}

func main() {
	flag := "--upper"
	args := os.Args[1:]

	convertToUpper := false
	for _, arg := range args {
		if arg == flag {
			convertToUpper = true
		} else {
			res := findChar(Atoi(arg))
			if convertToUpper {
				res = res - 'a' + 'A'
			}
			z01.PrintRune(res)
		}
	}
	z01.PrintRune('\n')
}
