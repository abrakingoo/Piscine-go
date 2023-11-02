package piscine

func ConcatParams(args []string) string {
	res := ""
	newline := "\n"

	for i, char := range args {
		res += char
		if i < len(args)-1 {
			res += newline
		}
	}

	return res
}
