func twoSum(nums []int, target int) []int {
	diffMap := make(map[int]int)
	for i, num := range nums {
		find := target - num
		if idx, ok := diffMap[find]; ok {
			return []int{idx, i}
		} else {
			diffMap[num] = i
		}
	}
	return []int{}
}
