func moveZeroes(nums []int) []int {
	w, r := 0, 0
	for w < len(nums) {
		if r == len(nums) {
			nums[w] = 0
			w++
			continue
		}
		if nums[r] != 0 {
			nums[w] = nums[r]
			w++
		}
		r++
	}
	return nums
}

