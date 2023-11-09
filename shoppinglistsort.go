package piscine

func ShoppingListSort(slice []string) []string {
	newList := []string{}

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	for i := 0; i < len(slice); i++ {
		newList = append(newList, slice[i])
	}
	return newList

}
