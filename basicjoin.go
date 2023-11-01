package piscine

func BasicJoin(elems []string) string {
	new := ""

	for i := 0; i < len(elems); i++ {
		new += elems[i]
	}
	return new
}
