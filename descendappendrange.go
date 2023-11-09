package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}

	result := make([]int, 0, max-min)
	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result
}
