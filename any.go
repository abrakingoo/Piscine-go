package piscine

func Any(f func(string) bool, a []string) bool {
	for _, char := range a {
		res := f(char)
		if res {
			return true
		}
	}
	return false
}
