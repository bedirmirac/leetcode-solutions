func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for key := range m {
		if m[key] == 1 {
			return key
		}
	}
	return 0
}

