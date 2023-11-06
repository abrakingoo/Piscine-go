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
	str := "x = " + convertToString(points.x) + ", y = " + convertToString(points.y)
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func convertToString(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	digits := ""
	for n > 0 {
		d := n % 10
		digits = string('0'+d) + digits
		n /= 10
	}
	return sign + digits
}
