func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	m2 := make(map[rune]int)
	var difference byte
	for _, letter := range s {
		m[letter]++
	}
	for _, d := range t {
		m2[d]++
	}
	for _, d := range t {
		val, ok := m[d]
		if !ok {
			difference = byte(d)
			break
		}
		if m2[d]-val != 0 {
			difference = byte(d)
		}
	}
	return difference
}

