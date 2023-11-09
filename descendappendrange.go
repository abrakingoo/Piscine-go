package piscine

func DescendAppendRange(min, max int) []int {
	if max > min {
		res := []int{}
		for i := max - min - 1; i >= 0; i-- {
			res[i] = max - i
		}
		return res
	}
	return []int{}
}
