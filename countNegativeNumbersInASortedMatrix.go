func countNegatives(grid [][]int) int {
	if grid == nil {
		return 0
	}
	m := len(grid[0]) - 1
	c := 0
	for i := 0; i < len(grid); i++ {
		if m >= 0 && grid[i][m] < 0 {
			c += len(grid) - i
			m--
			if m >= 0 && grid[i][m] < 0 {
				i--
			}
		} else {
			if i != len(grid)-1 && m >= 0 {
				if grid[i+1][m] < 0 {
					c += len(grid) - i - 1
					m--
				}
			}
		}
	}

	return c
}

