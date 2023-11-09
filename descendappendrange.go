package piscine

func DescendAppendRange(min, max int) []int {
	res := []int{}

	if max <= min {
		return []int{}
	}
	for i := max; i > min; i-- {
		res = append(res, i)
	}
	return res
}
