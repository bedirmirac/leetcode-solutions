func missingNumber(nums []int) int {
	missing, r, sum := 0, 0, 0
	exist := false
	for _, num := range nums {
		sum += num
		if num > r {
			r = num
		}
		if num == 0 {
			exist = true
		}
	}

	saved := r
	for ; r != 0; r-- {
		sum -= r
	}
	missing = sum * -1

	if !exist {
		return 0
	}

	if missing == 0 {
		missing = saved + 1
	}
	return missing
}

