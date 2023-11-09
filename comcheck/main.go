package main

import (
	"fmt"
	"os"
)

func main() {
	keyWord := []string{"01", "galaxy", "galaxy 01"}
	args := os.Args[1:]
	alertPrinted := false
	for _, char := range args {
		for _, key := range keyWord {
			if key == char && !alertPrinted {
				fmt.Println("Alert!!!")
				alertPrinted = true
				break
			}
		}
	}
}
