package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	isSorted := true
	for isSorted {
		isSorted = false
		current := l

		for current.Next != nil {
			if current.Data > current.Next.Data {
				temp := current.Data
				current.Data = current.Next.Data
				current.Next.Data = temp
				isSorted = true
			}
			current = current.Next
		}
	}

	return l
}
