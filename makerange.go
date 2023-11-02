package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	res := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		res[i] = min + i
	}
	return res
}
