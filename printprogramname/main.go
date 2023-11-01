package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args[0]
	rname := []rune(name)

	for i := 0; i < len(rname); i++ {
		z01.PrintRune(rname[i])
	}
	z01.PrintRune('\n')
}
