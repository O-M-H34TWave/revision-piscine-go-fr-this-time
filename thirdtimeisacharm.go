package piscine

func ThirdTimeIsACharm(str string) string {
	res := ""
	if str == "" {
		return "\n"
	}
	if len(str) < 2 {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	return res + "\n"
}
