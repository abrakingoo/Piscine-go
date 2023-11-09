package piscine

func StringToIntSlice(str string) []int {
	if str != "" {
		res := []int{}
		for _, char := range str {
			res = append(res, int(char))
		}
		return res
	}
	return []int(nil)
}
