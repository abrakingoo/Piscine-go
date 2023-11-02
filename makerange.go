package piscine

func MakeRange(min, max int) []int {
	len := max - min
	res := make([]int, len)

	if min >= max {
		return []int{}
	}
	for i := min; i < max; i++ {
		res[i] = min + i
	}
	return res
}
