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
	length := 0

	for divisor <= n/10 {
		divisor *= 10
		length++
	}
	length++

	for i := 0; i < length; i++ {
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
	count := 0
	for range list {
		count++
	}

	for i := 0; i < count; i++ {
		printValue(list[i])
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
