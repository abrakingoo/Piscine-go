package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a)-1; j++ {
			res := f(a[i], a[j])
			if res < 0 {
				return false
			}
		}
	}
	return true
}
