package piscine

func LastRune(s string) rune {
	Rune := []rune(s)

	count := 0
	for index := range s {
		count = index
	}

	return Rune[count]
}
