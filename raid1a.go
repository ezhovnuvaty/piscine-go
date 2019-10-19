package raid1a

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x < 0 || y < 0 {
	} else {
		if x == 0 || y == 0 {
		} else {
			if x == 1 && y == 1 {
				z01.PrintRune('o')
				z01.PrintRune(10)
			} else {
				if y == 1 {
					for i := 1; i <= x; i++ {
						if i == 1 || i == x {
							z01.PrintRune('o')
						} else {
							z01.PrintRune('-')
						}
					}
					z01.PrintRune(10)
				} else {
					for i := 1; i <= x; i++ {
						if i == 1 || i == x { 
							z01.PrintRune('o')
						} else {
							z01.PrintRune('-')
						}
					}
					z01.PrintRune(10)
					for j := 1; j <= y-2; j++ {
						for k := 1; k <= x; k++ {
							if k == 1 || k == x {
								z01.PrintRune('|')
							} else {
								z01.PrintRune(' ')
							}
						}
						z01.PrintRune(10)
					}
					for i := 1; i <= x; i++ {
						if i == 1 || i == x {
							z01.PrintRune('o')
						} else {
							z01.PrintRune('-')
						}
					}
					z01.PrintRune(10)
				}
			}
		}

	}
}
