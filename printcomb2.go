package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if i != j {
				if i < 10 {
					z01.PrintRune('0')
				}
				printNumber(i)
				z01.PrintRune(' ')
				if j < 10 {
					z01.PrintRune('0')
				}
				printNumber(j)

				if i != 98 || j != 99 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

