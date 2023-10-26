package piscine

func UltimateDivMod(a *int, b *int) {
	tempa := *a / *b
	tempb := *a % *b
	*a = tempa
	*b = tempb
}
