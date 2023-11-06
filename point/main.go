package main

import "github.com/01-edu/z01"

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func main() {
	points := &point{}
	setPoint(points)

	// Printing the values of x and y
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.y)
	z01.PrintRune('\n')
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	digits := []rune{}
	for n > 0 {
		digit := rune('0' + n%10)
		digits = append([]rune{digit}, digits...)
		n /= 10
	}
	if negative {
		z01.PrintRune('-')
	}
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}
