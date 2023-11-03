package piscine

func Split(s, sep string) []string {
	new := []string{}
	sRune := []rune(s)
	start := 0

	for i, _ := range sRune {
		if string(sRune[i:i+len(sep)]) == sep {
			new = append(new, string(sRune[start:i]))
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	if start < len(sRune) {
		new = append(new, string(sRune[start:]))
	}

	return new
}
