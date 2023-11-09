package piscine

func ActiveBits(n int) int {
	count := 0

	for i := 0; i < 32; i++ {
		count += n % 2
		n = n / 2
	}
	return count
}
