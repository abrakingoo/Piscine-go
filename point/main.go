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

func printValue(n int) {
	divisor := 1
	for divisor <= n/10 {
		divisor *= 10
	}
	for divisor > 0 {
		digit := n / divisor
		z01.PrintRune('0' + digit)
		n %= divisor
		divisor /= 10
	}
}

func main() {
	points := &point{}
	setPoint(points)
	list := []int{points.y, points.x}

	for i := 0; i < len(list); i++ {
		printValue(list[i])
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
