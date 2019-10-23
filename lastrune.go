package piscine

func LastRune(s string) rune {
	result := []rune(s)

	for index, value := range result {
		if index == len(result)-1 {
			return value
		}
	}
	return 0
}
