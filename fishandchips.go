package piscine

func FishAndChips(n int) string {
	res := ""

	if n < 0 {
		res = "error: number is negative"
	}
	if n%2 != 0 || n%3 != 0 {
		res = "error: non divisible"
	}
	if n%2 == 0 {
		res = "fish"
	}
	if n%3 == 0 {
		res = "chips"
	}
	if n%2 == 0 && n%3 == 0 {
		res = "fish and chips"
	}

	return res
}
