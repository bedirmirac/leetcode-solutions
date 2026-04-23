func mirrorDistance(n int) int {
	reverse, remainder := 0, 0
	original := n
	distance := 0
	for n != 0 {
		remainder = n % 10
		reverse = reverse*10 + remainder
		n /= 10
	}
	distance = original - reverse
	if distance < 0 {
		return distance - 2*distance
	}
	return distance
}

