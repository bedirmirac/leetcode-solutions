func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		current := s[i]
		if i == len(s)-1 {
			sum += m[current]
			break
		}
		if i+1 <= len(s)-1 {
			next := s[i+1]
			if m[current] < m[next] {
				sum -= m[current]
			} else {
				sum += m[current]
			}
		}
	}
	return sum
}


