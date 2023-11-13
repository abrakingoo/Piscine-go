package piscine

func ListSize(l *List) int {
	count := 0
	Head := l.Head

	for Head != nil {
		count++
		Head = Head.Next
	}
	return count
}
