package piscine

func AppendRange(min, max int) []int {

	var result []int

	if min < max || min != max {

		for i := min; i < max; i++ {
			result = append(result, i)
		}
		return result
	}
	return nil
}
