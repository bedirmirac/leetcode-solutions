func findWords(words []string) []string {
	row1 := map[byte]int{
		'q': 0, 'w': 0, 'e': 0, 'r': 0, 't': 0, 'y': 0, 'u': 0, 'i': 0, 'o': 0, 'p': 0,
		'Q': 0, 'W': 0, 'E': 0, 'R': 0, 'T': 0, 'Y': 0, 'U': 0, 'I': 0, 'O': 0, 'P': 0,
	}

	row2 := map[byte]int{
		'a': 0, 's': 0, 'd': 0, 'f': 0, 'g': 0, 'h': 0, 'j': 0, 'k': 0, 'l': 0,
		'A': 0, 'S': 0, 'D': 0, 'F': 0, 'G': 0, 'H': 0, 'J': 0, 'K': 0, 'L': 0,
	}

	row3 := map[byte]int{
		'z': 0, 'x': 0, 'c': 0, 'v': 0, 'b': 0, 'n': 0, 'm': 0,
		'Z': 0, 'X': 0, 'C': 0, 'V': 0, 'B': 0, 'N': 0, 'M': 0,
	}

	newSlice := []string{}

	for _, word := range words {
		r1 := 0
		r2 := 0
		r3 := 0
		for i := 0; i < len(word); i++ {
			if _, ok := row1[word[i]]; ok {
				r1++
			} else if _, ok := row2[word[i]]; ok {
				r2++
			} else if _, ok := row3[word[i]]; ok {
				r3++
			} else {
				r1 = 1
				r2 = 1
				break
			}
		}
		if !((r1 != 0 && r2+r3 != 0) || (r2 != 0 && r1+r3 != 0) || (r3 != 0 && r1+r2 != 0)) {
			newSlice = append(newSlice, word)
		}
	}

	return newSlice
}

