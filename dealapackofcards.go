package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	val := 0
	for i := 1; i <= len(deck)/3; i++ {
		fmt.Printf("Player ")
		fmt.Printf("%v", i)
		fmt.Printf(": ")
		for j := val; j < val+3; j++ {
			fmt.Printf("%v", deck[j])
			if j < val+2 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		val += 3
	}
}
