package piscine

func IsPrintable(str string) bool {
	for _, v := range []rune(str) {
		if v >= 0 && v <= 31 {
			return false
		}
	}
	return true
}
