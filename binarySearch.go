func search(nums []int, target int) int {
	idx := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		current := nums[mid]
		if current > target {
			right = mid - 1
		} else if current < target {
			left = mid + 1
		} else if current == target {
			idx = mid
			break
		}
	}
	return idx
}

