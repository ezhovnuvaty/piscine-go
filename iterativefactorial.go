package piscine

func iterativefactorial(nb int) int {
	result := 1
	if nb < 0 || nb >= 20 {
		return 0
	}
	for i := 0; i <= nb; i++ {
		result = result * i
	}
	return result
}