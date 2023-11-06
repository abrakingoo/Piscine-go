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
	printNumber(points.x)
	z01.PrintRune(' ')
	printNumber(points.y)
	z01.PrintRune('\n')
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	digits := []rune{}
	for n > 0 {
		d := n % 10
		digits = append([]rune{rune('0' + d)}, digits...)
		n /= 10
	}
	for i := 0; i < len(digits); i++ {
		z01.PrintRune(rune(digits[i]))
	}
}
