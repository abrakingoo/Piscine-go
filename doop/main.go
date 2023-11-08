package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	maxInt64Str = "9223372036854775807"  // Maximum int64 value as a string
	minInt64Str = "-9223372036854775808" // Minimum int64 value as a string
)

func IsOverflow(s string) bool {
	if len(s) < len(maxInt64Str) {
		return false
	} else if len(s) > len(maxInt64Str) {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i] < maxInt64Str[i] {
			return true
		}
	}
	return false
}

func IsUnderflow(s string) bool {
	if len(s) < len(minInt64Str) {
		return false
	} else if len(s) > len(minInt64Str) {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i] < minInt64Str[i] {
			return true
		}
	}
	return false
}

func Isdigit(n byte) bool {
	return n >= '0' && n <= '9'
}

func Atoi(s string) int {
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0 // indicate invalid input
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}



func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		for _, c := range "9223372036854775808" {
			z01.PrintRune(c)
		}
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		return
	}

	num1 := args[0]
	num2 := args[2]
	oper := args[1]

	if IsUnderflow(num1) || IsOverflow(num1) || IsUnderflow(num2) || IsOverflow(num2) {
		return
	}

	for i := 0; i < len(num1); i++ {
		if !Isdigit(num1[i]) {
			return
		}
	}

	for i := 0; i < len(num2); i++ {
		if !Isdigit(num2[i]) {
			return
		}
	}

	operators := []string{"+", "-", "/", "%", "*"}

	validOperator := false
	for i := 0; i < len(operators); i++ {
		if oper == operators[i] {
			validOperator = true
			break
		}
	}

	if !validOperator {
		return
	}

	val1 := Atoi(num1)
	val2 := Atoi(num2)

	res := 0

	switch oper {
	case "+":
		if (val1 > 0) && (val2 > 9223372036854775807-val1) {
			return
		}
		if (val1 < 0) && (val2 < -9223372036854775808-val1) {
			return
		}
		res = val1 + val2

	case "-":
		if (val1 > 0) && (val2 > 9223372036854775807-val1) {
			return
		}
		if (val1 < 0) && (val2 < -9223372036854775808-val1) {
			return
		}
		res = val1 - val2

	case "/":
		if (val1 > 0) && (val2 > 9223372036854775807-val1) {
			return
		}
		if (val1 < 0) && (val2 < -9223372036854775808-val1) {
			return
		}
		if val2 == 0 {
			printstr("No division by 0")
			z01.PrintRune('\n')
			return
		}
		res = val1 / val2

	case "%":
		if (val1 > 0) && (val2 > 9223372036854775807-val1) {
			return
		}
		if (val1 < 0) && (val2 < -9223372036854775808-val1) {
			return
		}
		if val2 == 0 {
			printstr("No modulo by 0")
			z01.PrintRune('\n')
			return
		}
		res = val1 % val2

	case "*":
		res = val1 * val2
	}

	PrintNbr(res)
	z01.PrintRune('\n')
}
