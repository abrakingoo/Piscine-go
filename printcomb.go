package piscine

import "fmt"

func PrintComb() {
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				fmt.Print(i, j, k)
				if i < 7 {
					fmt.Print(", ")
				} else if i == 7 && j < 8 {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}
