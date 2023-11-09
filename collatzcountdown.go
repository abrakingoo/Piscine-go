package piscine

func CollatzCountdown(start int) int {
	count := 1

	for start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = 3*start + 1
		}
		count++
	}
	return count - 1
}
