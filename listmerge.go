package piscine

func ListMerge(l1 *List, l2 *List) {
	current := l1.Head

	for current != nil {
		if current.Next == nil {
			current.Next = l2.Head
			break
		}
		current = current.Next
	}
}
