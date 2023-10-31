package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	count := 0
	maxDigit := 0

	if n == 0 {
		z01.PrintRune('0')
	}

	// Find the maximum digit in n
	temp := n
	for temp != 0 {
		digit := temp % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		temp /= 10
	}

	// Create the list slice based on the maximum digit
	list := make([]int, maxDigit+1)

	// Count the occurrences of each digit in n
	for n != 0 {
		digit := n % 10
		list[digit]++
		n /= 10
		count++
	}

	// Print the digits in ascending order
	for i := 0; i <= maxDigit; i++ {
		for j := 0; j < list[i]; j++ {
			z01.PrintRune(rune(i + '0'))
		}
	}
}
