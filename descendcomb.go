package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for x := '9'; x >= '0'; x-- {
				for y := '9'; y >= '0'; y-- {
					if i > x || (i == x && j >= y) {
						if !(i == '9' && j == '8' && x == '9' && y == '7') {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(' ')
							z01.PrintRune(x)
							z01.PrintRune(y)
							if !(i == '9' && j == '0' && x == '0' && y == '0') {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
