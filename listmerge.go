package piscine

func ListMerge(l1 *List, l2 *List) {
	current := l1.Head

	if current == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	}
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
