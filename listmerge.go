package piscine

func ListMerge(l1 *List, l2 *List) {
	current := l1.Head

	if current == nil {
		l1.Head = l2.Head
	}

	for current.Next != nil {
		current = current.Next
	}
	current.Next = l2.Head
}
