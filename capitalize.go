package piscine

func Capitalize(s string) string {
	res := ""
	flag := true
	for _, r := range s {
		if flag {
			if r >= 'a' && r <= 'z' {
				res = res + string(r+('A'-'a'))
				flag = false
			} else if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') {
				res = res + string(r)
				flag = false
			} else {
				res = res + string(r)
				flag = true
			}
		} else {
			if r >= 'A' && r <= 'Z' {
				res = res + string(r-('A'-'a'))
			} else if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') {
				res = res + string(r)
			} else {
				res = res + string(r)
				flag = true
			}
		}
	}
	return res
}
