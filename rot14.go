package piscine

func Rot14(s string) string {
	rot14 := ""
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= 'A' && s[i] <= 'Z':
			res := (s[i]-'A'+14)%26 + 'A'
			rot14 += string(res)
		case s[i] >= 'a' && s[i] <= 'z':
			res := (s[i]-'a'+14)%26 + 'a'
			rot14 += string(res)
		default:
			rot14 += string(s[i])
		}
	}
	return rot14
}
