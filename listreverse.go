package piscine

func ListReverse(l *List) {
	current := l.Head
	var prev *NodeL = nil

	for current != nil {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}

	l.Head = prev
}
