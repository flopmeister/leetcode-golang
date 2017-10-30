import "unicode"

func myAtoi(str string) int {
	result := 0
	int_max := 2147483647
	int_min := -2147483648
	// Trim all the empty spaces
	runes := []rune(str)
	i := 0
	for ; i < len(runes); i++ {
		if runes[i] != ' ' {
			break
		}
	}
	// if result is empty , return
	if i == len(runes) {
		return result
	}

	sign := 1
	if runes[i] == '+' {
		i++
	} else if runes[i] == '-' {
		sign = -1
		i++
	}

	for ; i < len(runes); i++ {
		if !unicode.IsNumber(runes[i]) {
			break
		}
		if result > int_max/10 || (result == int_max/10 && int(runes[i]-'0') > 7) {
			if sign == 1 {
				return int_max
			} else {
				return int_min
			}
		}
		result = result*10 + int(runes[i]-'0')
	}
	return result * sign
}