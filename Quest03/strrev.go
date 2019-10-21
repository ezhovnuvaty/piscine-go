package piscine

func StrRev(s string) string {
	var arr = []rune(s)
	var len int = 0
	for a := range arr {
		len = a
	}
	for a, b := 0, len; a < b; a, b = a+1, b-1 {
		arr[a], arr[b] = arr[b], arr[a]
	}
	return string(arr)
}
