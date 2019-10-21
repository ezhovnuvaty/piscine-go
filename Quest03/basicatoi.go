package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, value := range s {
		count := 0
		if value < '0' || value > '9'{
			return 0
		}
		for i := '1'; i <= value; i++ {
			count++
		}
		result = result*10 + count
	}
	return result
}
