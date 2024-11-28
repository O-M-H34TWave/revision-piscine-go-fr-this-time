package piscine

func FromTo(from int, to int) string {
	res := ""
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	if from < to {
		for i := from; i <= to; i++ {
			res += itoa(i)
			if i != to {
				res += ", "
			}
		}
	}
	if from > to {
		for i := from; i >= to; i-- {
			res += itoa(i)
			if i != to {
				res += ", "
			}
		}
	}

	return res + "\n"
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := ""
	if n < 0 {
		n = -n
		neg = "-"
	}
	res := ""
	for n > 0 {
		digit := n % 10
		res = string(rune('0'+digit)) + res
		n /= 10
	}
	return neg + res
}
