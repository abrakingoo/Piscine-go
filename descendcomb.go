package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '9'; i >= '0'; i++ {
		for j := '9'; j >= '0'; j++ {
			if i < j {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(i)
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')
}
