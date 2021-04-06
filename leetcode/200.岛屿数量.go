package leetcode

/*
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/
func numIslands(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	allIsland := 0
	var dfs func(grid [][]byte, r int, c int)
	dfs = func(grid [][]byte, r int, c int) {
		grid[r][c] = '0'
		if (r-1) >= 0 && grid[r-1][c] == '1' {
			dfs(grid, r-1, c)
		}
		if (r+1) < nr && grid[r+1][c] == '1' {
			dfs(grid, r+1, c)
		}
		if (c-1) >= 0 && grid[r][c-1] == '1' {
			dfs(grid, r, c-1)
		}
		if (c+1) < nc && grid[r][c+1] == '1' {
			dfs(grid, r, c+1)
		}
	}
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				allIsland++
				dfs(grid, r, c)
			}
		}
	}
	return allIsland
}
