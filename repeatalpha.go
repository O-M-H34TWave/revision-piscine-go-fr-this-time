package piscine

func RepeatAlpha(s string) string {
	res := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			count = int(s[i]) - 64
			for count > 0 {
				res += string(s[i])
				count--
			}
		}
		if s[i] >= 97 && s[i] <= 122 {
			count = int(s[i]) - 96
			for count > 0 {
				res += string(s[i])
				count--
			}
		}
	}
	return res
}
