package piscine

func ListSize(l *List) int {
	count := 0
	Head := l.Head

	if Head == nil {
		return 0
	}
	for l.Head != nil {
		count++
		Head = Head.Next
	}
	return count
}
