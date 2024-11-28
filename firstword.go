package piscine

func FirstWord(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && len(s) > 0 {
			break
		}
		if s[i] != ' ' {
			res += string(s[i])
		}
	}
	return res + "\n"
}
