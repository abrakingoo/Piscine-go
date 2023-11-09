package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)
	res := make([][]string, length)
	for i := length - 1; i >= 0; i-- {
		res[length-1-i] = podium[i]
	}
	return res
}
