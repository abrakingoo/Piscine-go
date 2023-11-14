package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	count := 0

	for current != nil && count <= pos {
		if count == pos {
			return current
		}
		current = current.Next
	}

	return nil
}
