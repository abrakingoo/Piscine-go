package main

import "fmt"

func main() {

	// This is to print a to z using ascii code
	for char := 97; char <= 122; char++ {
		fmt.Printf("%c", char)
	}
	fmt.Println("")

}
