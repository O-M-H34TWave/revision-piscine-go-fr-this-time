package piscine

func CamelToSnakeCase(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 && i != 0 {
			res += "_" + string(s[i])
		}else{
			res+=string(s[i])
		}

	}
	return res
}
