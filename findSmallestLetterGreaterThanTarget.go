func check(current, next byte) bool {
	return current != next
}

func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1
	counter := -1
	for low <= high {
		mid := low + (high-low)/2
		current := letters[mid]
		if target == current {
			counter++
			current = letters[mid+counter]
			if mid+counter+1 > len(letters)-1 {
				return letters[0]
			}
			if check(current, letters[mid+1+counter]) {
				return letters[mid+counter+1]
			}
		}
		if current > target {
			high = mid - 1
		} else if current < target {
			if current < target {
				low = mid + 1
			}
		}
	}
	if low > len(letters)-1 {
		return letters[0]
	}
	return letters[low]
}

