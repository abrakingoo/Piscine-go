package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func printNumber(n int) []rune {
	digits := []rune{}
	for n > 0 {
		d := n % 10
		digits = append([]rune{rune('0' + d)}, digits...)
		n /= 10
	}

	return digits
}
func main() {
	points := &point{}

	setPoint(points)
	x := printNumber(points.x)
	y := printNumber(points.y)
	str := "x = " + string(x) + ", " + "y = " + string(y)
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
