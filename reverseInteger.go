import "math"

func reverse(x int) int {
	var remainder int = 0
	var reverse int = 0

	for x != 0 {
		remainder = x % 10
		x /= 10
		reverse = reverse*10 + remainder
	}
	if math.Pow(-2, 31) < float64(reverse) && float64(reverse) < math.Pow(2, 31)-1.0 {
		return reverse
	}
	return 0
}
