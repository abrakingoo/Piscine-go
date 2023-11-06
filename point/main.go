package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digit := n % 10
	rest := n / 10

	if rest != 0 {
		printNumber(rest)
	}
	z01.PrintRune('0' + rune(digit))
}

func main() {
	points := &point{}

	setPoint(points)
	list := []int{points.x, points.y}

	for i := 0; i < len(list); i++ {
		printNumber(list[i])
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
