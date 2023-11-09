package piscine

func PodiumPosition(podium [][]string) [][]string {
	res := [][]string{}
	for i := len(podium) - 1; i >= 0; i-- {
		res = append(res, podium[i])
	}
	return res
}
