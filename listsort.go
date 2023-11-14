package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	current := l

	for current != nil {
		if current.Data > current.Next.Data {
			temp := current.Data
			current.Data = current.Next.Data
			current.Next.Data = temp
		}
		current = current.Next
	}
	return l
}
