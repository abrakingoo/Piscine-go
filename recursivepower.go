package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	if nb < 0 {
		if power%2 == 0 {
			return RecursivePower(-nb, power)
		} else {
			return -RecursivePower(-nb, power)
		}
	}
	return nb * RecursivePower(nb, power-1)
}
