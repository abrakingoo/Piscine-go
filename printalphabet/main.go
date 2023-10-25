package main

import "github.com/01-edu/z01"

func main() {

	// This is to print a to z using ascii code
	for char := 'a'; char <= 'z'; char++ {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
