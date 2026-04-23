func plusOne(digit []int) []int {
	lastDigit := digit[len(digit)-1]
	c := 0
	if lastDigit == 9 {
		for i := len(digit) - 1; i >= 0; i-- {
			if digit[i] == 9 {
				c++
				continue
			}
			break
		}
		if len(digit)-c-1 > -1 {
			lastDigit = digit[len(digit)-c-1] + 1
			digit = digit[:len(digit)-c-1]
		} else {
			lastDigit = 1
			digit = digit[:len(digit)-c]

		}
		digit = append(digit, lastDigit)
		for c != 0 {
			digit = append(digit, 0)
			c--
		}
		return digit
	}
	lastDigit += 1
	digit = digit[:len(digit)-1] // remove the last digit
	digit = append(digit, lastDigit)
	return digit
}
