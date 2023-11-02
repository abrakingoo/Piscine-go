package piscine

func MakeRange(min, max int) []int {
	res := make([]int, max)

	if min >= max {
		return res
	}
	for i := min; i < max; i++ {
		res[i] = i
	}
	return res
}
