package piscine

func CheckNumber(arg string) bool {
	for _, v := range arg {
		if v>='0' && v<='9'{
            return true
        }
	}
	return false
}
