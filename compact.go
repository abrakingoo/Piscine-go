package piscine

func Compact(ptr *[]string) int {
	count := 0
	newS := make([]string, 0)
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			newS = append(newS, (*ptr)[i])
			count++
		}
	}
	*ptr = newS
	return count
}
