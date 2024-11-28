package piscine

func LastWord(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && len(res)> 0{
			break
		}
		if s[i] != ' ' {
			res = string(s[i]) + res
		}
	}
	return res + "\n"
}
