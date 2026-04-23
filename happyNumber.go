func isHappy(n int) bool {
	m := map[int]int{}
	sum := 0
	for {
		ld := n % 10
		n /= 10
		square := ld * ld
		sum += square
		if n == 0 && sum != 0 {
			m[sum]++
			n = sum

			if sum == 1 {
				return true
			}
			sum = 0
			for _, val := range m {
				if val > 1 {
					return false
				}
			}
		}

	}
}
