package piscine

func ShoppingListSort(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}
