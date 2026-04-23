func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	m2 := make(map[byte]bool)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		current := s[i]
		target := t[i]
		val, ok := m[current]
		if ok && val != target {
			return false
		}
		if m2[target] && val != target {
			return false
		}
		if !ok {
			m[current] = t[i]
			m2[t[i]] = true
		}

	}
	return true
}

