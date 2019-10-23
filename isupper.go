package piscine

func IsUpper(str string) bool {
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
		} else {
			return false
		}
	}
	return true
}
