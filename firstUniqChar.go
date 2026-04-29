func firstUniqChar(s string) int {
	pairs := make(map[rune]int)
	for _, letter := range s {
		pairs[letter]++
	}
	for i, l := range s {
		if pairs[l] == 1 {
			return i
		}
	}

	return -1
}
