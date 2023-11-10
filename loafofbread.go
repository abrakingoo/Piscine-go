package piscine

import (
	"github.com/01-edu/z01"
)

func LoafOfBread(str string) string {
	length := 5
	if len(str) < length {
		return "Invalid Output\n"
	}
	result := ""
	for i, char := range str {
		if i%6 != 5 {
			result += string(char)
		}
	}
	return result
}
