package piscine

func Index(s string, toFind string) int {
	find := toFind[0]
	for i := 0; i < len(s); i++ {
		if s[i] == find {
			return i
		}
	}
	return -1
}
