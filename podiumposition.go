package piscine

func PodiumPosition(podium [][]string) [][]string {
	res := podium

	for j := 0; j < len(res); j++ {
		for i := 0; i < len(res)-1; i++ {
			if res[i][0] > res[i+1][0] {
				res[i], res[i+1] = res[i+1], res[i]
			}
		}
	}
	return res
}
