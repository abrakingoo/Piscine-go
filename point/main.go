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

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}
	str := ""
	for n > 0 {
		digit := n % 10
		str = string('0'+digit) + str
		n /= 10
	}
	if negative {
		str = "-" + str
	}
	return str
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)
	xStr := "x = " + intToString(points.x) + ", "
	yStr := "y = " + intToString(points.y)

	// Printing the values of x and y
	printStr(xStr)
	printStr(yStr)
	printStr("\n")
}
