package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for x := 2; x <= power; x++ {
		result = result * nb
	}
	return result
}
