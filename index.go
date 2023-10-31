package piscine

func Index(s string, toFind string) int {
	find := toFind[0]
	for index, char := range s {
		if s[char] == find {
			return index
		}
	}
	return 0
}
