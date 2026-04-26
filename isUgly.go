func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	primes := []int{2, 3, 5}
	for _, prime := range primes {
		for n%prime == 0 {
			n /= prime
		}
	}

	if n == 1 {
		return true
	} else {
		return false
	}
}

