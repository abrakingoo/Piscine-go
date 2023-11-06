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
	printChar('x')
	printChar(' ')
	printChar('=')
	printChar(' ')
	printNumber(points.x)
	printChar(',')
	printChar(' ')
	printChar('y')
	printChar(' ')
	printChar('=')
	printChar(' ')
	printNumber(points.y)
	printChar('\n')
}

func printChar(char rune) {
	z01.PrintRune(char)
}

func printNumber(number int) {
	if number == 0 {
		printChar('0')
		return
	}
	if number < 0 {
		printChar('-')
		number *= -1
	}
	digits := []rune{}
	for number > 0 {
		digit := '0' + rune(number%10)
		digits = append([]rune{digit}, digits...)
		number /= 10
	}
	for _, digit := range digits {
		printChar(digit)
	}
}
