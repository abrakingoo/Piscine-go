package piscine

func ListLast(l *List) interface{} {
	if l.Head != nil {
		return l.Head.Data
	}
	return nil
}
