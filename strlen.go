 
package piscine

func StrLen(str string) int {
	var arr = []rune(str)
	var len int = 0
	for a := range arr {
		len = a + 1
	}
	return len
}
