package piscine

import "github.com/01-edu/z01"

func PrintCombinations(n int) {
	if n <= 0 || n >= 10 {
		PrintStr("n should be between 0 and 10\n")
		return
	}
	digits := make([]rune, n)
	helper(&digits, 0, '0', n)
}

func helper(digits *[]rune, index int, startChar rune, n int) {
	if index == n {
		for i, v := range *digits {
			if i > 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return
	}
	for i := startChar; i < '9'+1; i++ {
		if contains(*digits, i) {
			continue
		}
		(*digits)[index] = i
		helper(digits, index+1, i+1, n)
	}
}
