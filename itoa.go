package piscine

func Itoa(n int) string {
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

/*
example :
Let’s use n = -456:

Is it 0? No.
Is it negative? Yes, save "-" and flip 456 to positive.
Start looping:
	First loop: 456 % 10 = 6 → q = "6".
	Second loop: 45 % 10 = 5 → q = "56".
	Third loop: 4 % 10 = 4 → q = "456".
	End the loop: n is now 0.
Add the sign: Combine "-" and "456".
	Result: Return "-456".*/
