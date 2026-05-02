func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return false
	}
	dx := coordinates[1][0] - coordinates[0][0]
	dy := coordinates[1][1] - coordinates[0][1]

	for i := 2; i < len(coordinates); i++ {
		currDx := coordinates[i][0] - coordinates[0][0]
		currDy := coordinates[i][1] - coordinates[0][1]
		if dy*currDx != dx*currDy {
			return false
		}
	}
	return true
}

