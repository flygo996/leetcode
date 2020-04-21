package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var sum = 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				sum++
				dfs(grid, i, j)
			}
		}
	}
	return sum
}

func dfs(grid [][]byte, r, c int) {
	var nr = len(grid)
	var nc = len(grid[0])
	if r < 0 || c < 0 || r > nr-1 || c > nc-1 || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}