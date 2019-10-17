package piscine

func UltimadeDivMod(a *int, b *int) {
	a2 := *a
	b2 := *b
	*a = *a / *b
	*b = a2 % b2
}
