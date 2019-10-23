package piscine

func Index(s string, toFind string) int {

	j := []rune(s)
	l := []rune(toFind)
	n := arrLength(j)
	k := arrLength(l)
	index := -1

	if arrLength(j) < arrLength(l) || arrLength(j) > 13 {
		return index
	} else if arrLength(j) == 0 {
		return index
	} else if arrLength(l) == 0 {
		return 0
	}
	for i := 0; i <= n-k; i++ {
		if toFind == s[i:i+k] {
			return (i)
		}
	}
	return -1
}

func arrLength(arr []rune) int {
	ln := 0
	for range arr {
		ln++
	}
	return ln

}
