package piscine

func Map(f func(int) bool, a []int) []bool {
	res := []bool{}
	for _, char := range a {
		boolean := f(char)
		res = append(res, boolean)
	}
	return res
}
