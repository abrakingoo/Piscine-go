package piscine

import (
	"github.com/01-edu/z01"
)

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var result string
	counter := 0
	for _, char := range str {
		if counter < 5 {
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
