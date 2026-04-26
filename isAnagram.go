func isAnagram(s string, t string) bool {
	if len(t) > len(s) {
		return false
	}
	sLetters := make(map[byte]int)
	for _, letter := range s {
		sLetters[byte(letter)]++
	}

	for _, letter := range t {
		_, ok := sLetters[byte(letter)]
		if ok {
			sLetters[byte(letter)]--
		}
	}

	for _, value := range sLetters {
		if value != 0 {
			return false
		}
	}

	return true
}

