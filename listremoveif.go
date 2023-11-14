package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	var prev *NodeL = nil

	for current != nil {
		nextNode := current.Next
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}
			current = nil
		} else {
			prev = current
		}
		current = nextNode
	}
}
