 
package piscine

func StrLen(str string) int {
	var array = []rune(str)
	var len int = 0
	for a := range array {
		len = a + 1
	}
	return len
}
