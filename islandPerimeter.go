func check(grid [][]int, r, c int) int {
	perimeter := 4
	if r-1 >= 0 && grid[r-1][c] == 1 {
		perimeter--
	}
	if r+1 < len(grid) && grid[r+1][c] == 1 {
		perimeter--
	}
	if c-1 >= 0 && grid[r][c-1] == 1 {
		perimeter--
	}
	if c+1 < len(grid[0]) && grid[r][c+1] == 1 {
		perimeter--
	}

	return perimeter
}

func islandPerimeter(grid [][]int) int {
	totalPerimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				totalPerimeter += check(grid, i, j)
			}
		}
	}
	return totalPerimeter
}

