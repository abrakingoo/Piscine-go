package piscine

import (
	"github.com/01-edu/z01"
)

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var result string
	var count int
	for _, char := range str {
		if count < 5 {
			if char != ' ' {
				z01.PrintRune(char)
				result += string(char)
				count++
			}
		} else {
			z01.PrintRune('\n')
			break
		}
	}
	return result
}
