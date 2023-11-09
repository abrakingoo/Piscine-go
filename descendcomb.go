package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for x := '9'; x >= '0'; x-- {
				for y := '8'; y <= '0'; y-- {
					if i < x || (i == x && j < y) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(x)
						z01.PrintRune(y)
						if !(i == '0' && j == '1' && x == '0' && y == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
