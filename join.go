package piscine

func Join(strs []string, sep string) string {
	var res string = strs[0]
	for i := 1; i <= len(strs[1:]); i++ {
		res = res + sep + string(strs[i])
	}
	return res
}
