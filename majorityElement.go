func majorityElement(nums []int) int {
	m := make(map[int]int)
	target := len(nums) / 2
	for _, num := range nums {
		if val, ok := m[num]; ok {
			val++
			m[num] = val
		} else {
			m[num] = 1
		}
	}
	// compare
	temp := 0
	element := 0
	for key, val := range m {
		if val > target && val > temp {
			temp = val
			element = key
		}
	}
	return element
}


