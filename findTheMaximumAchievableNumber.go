func theMaximumAchievableX(num int, t int) int {
	if t == 0 {
		return num
	} else if t < 0 {
		return num + (2 * -t)
	}
	return num + (2 * t)
}

