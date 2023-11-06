package main

import "fmt"

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

	setPoint(*&points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
