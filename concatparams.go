package piscine

func ConcatParams(args []string) string {
	answer := ""
	for i, c := range args {
		if i == 0 {
			answer = c
			continue
		}
		answer = answer + "\n" + c
	}
	return answer
}
