package piscine

import (
	"github.com/01-edu/z01"
)

func LoafOfBread(str string) string {
	length := 5
	if len(str) < length {
		return "Invalid Output\n"
	}
	var result string
	counter := 0
	for _, char := range str {
		if counter < length {
			if char != ' ' {
				z01.PrintRune(char)
				result += string(char)
				counter++
			}
		} else {
			z01.PrintRune(' ')
			z01.PrintRune(char)
			result += " " + string(char)
			counter++
		}
	}
	z01.PrintRune('\n')
	return result + "\n"
}
