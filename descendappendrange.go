package piscine

func DescendAppendRange(min, max int) []int {
	res := []int{}

	if min >= max {
		return []int{}
	}
	for i := max; i < min; i-- {
		res = append(res, i)
	}
	return res
}
