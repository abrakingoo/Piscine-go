package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 48 && s[i] <= 57)) {
			return false
		}
	}
	return true
}
