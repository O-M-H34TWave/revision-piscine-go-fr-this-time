package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := 1; j <= 8; j++ {
			for k := 2; k <= 9; k++ {
				z01.PrintRune('0'+rune(i))
				z01.PrintRune('0'+rune(j))
				z01.PrintRune('0'+rune(k))
				if i!= 7 || j !=8 || k != 9{
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
