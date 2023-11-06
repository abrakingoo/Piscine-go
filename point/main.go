package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func main() {
	points := &point{}
	points.x = 42
	points.y = 21

	// Printing the values of x and y
	printStr("x = ")
	printInt(points.x)
	printStr(", y = ")
	printInt(points.y)
	printStr("\n")
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var printHelper func(int)
	printHelper = func(n int) {
		if n == 0 {
			return
		}
		printHelper(n / 10)
		z01.PrintRune('0' + rune(n%10))
	}
	printHelper(n)
}
