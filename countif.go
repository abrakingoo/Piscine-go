package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	for _, char := range tab {
		res := f(char)
		if res {
			count++
		}
	}
	return count
}
