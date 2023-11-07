package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		res := f(a[i], a[i+1])
		if res > 0 {
			descending = false
		} else if res < 0 {
			ascending = false
		}
	}

	return ascending || descending
}
