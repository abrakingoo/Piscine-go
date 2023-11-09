package piscine

func DescendAppendRange(min, max int) []int {
	res := []int{}

	if max > min {
		for i := max; i > min; i-- {
			res = append(res, i)
		}
	}
	return res
}
