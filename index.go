package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	find := toFind[0]
	if find >= 48 && find <= 57 {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == find {
			return i
		}
	}
	return -1
}
