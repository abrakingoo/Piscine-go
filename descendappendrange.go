package piscine

func DescendAppendRange(min, max int) []int {
	if max > min {
		res := []int{}
		for i := max - min - 1; i >= 0; i-- {
			res = append(res, max-i)
		}
		return res
	}
	return []int{}
}
