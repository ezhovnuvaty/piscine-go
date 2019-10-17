package piscine

func Atoi(s string) int {
	result := 0
	isNegative := 0
	for i, value := range s {

		if i == 0 && (value == '+' || value == '-') {
			if value == '-' {
				isNegative = 1
			}
			continue
		}
		count := 0
		if value < '0' || value > '9' {
			return 0
		}
		for i := '1'; i <= value; i++ {
			count++
		}
		result = result*10 + count
	}
	if isNegative == 1 {
		result = result * (-1)
	}
	return result
}
