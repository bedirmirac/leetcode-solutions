func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	remainder := 0
	reverse := 0
	for x != 0 {
		remainder = x % 10
		reverse = reverse*10 + remainder
		x /= 10
	}

	if original == reverse {
		return true
	}
	return false
}
