package main

import "fmt"

func main() {
	// This is to print a to z using ascii code
	for char := 97; char <= 122; char++ {
		fmt.Printf("%c", char)
	}
	fmt.Println("")

	// Second example to print a to z using character
	for char := 'a'; char <= 'z'; char++ {
		fmt.Printf("%c", char)
	}

}
