package piscine

func ToUpper(s string) string {

	var st string
	for _, i := range s {

		if i >= 'A' && i <= 'Z' {

			st += string(i + 32)

		} else {

			st += string(i)
		}
	}

	return st

}
