package piscine

func ListClear(l *List) {
	current := l.Head

	for current != nil {
		nextNode := current.Next
		current.Next = nil
		current = nextNode
	}

	l.Head = nil
	l.Tail = nil
}
