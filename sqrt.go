func mySqrt(x int) int {
	counter := 0
	sub := 1
	for x >= 0 {
		x -= sub
		sub += 2
		counter++
	}
	if x < 0 {
		counter--
	}
	return counter
}

