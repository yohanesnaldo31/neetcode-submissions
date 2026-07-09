func maxAreaOfIsland(grid [][]int) int {
    lengthRow := len(grid)
	lengthCol := len(grid[0])

	var dfs func(idxRow,idxCol int) int
	dfs = func(idxRow,idxCol int) int{
		if idxRow<0|| idxCol<0||idxRow>=lengthRow||idxCol>=lengthCol||grid[idxRow][idxCol]!=1{
			return 0
		}
		grid[idxRow][idxCol] = 2
		return 1 + dfs(idxRow+1,idxCol) + dfs(idxRow,idxCol+1) + dfs(idxRow-1,idxCol) + dfs(idxRow,idxCol-1)
	}

	maxCount := 0

	for row := 0; row<lengthRow ; row++ {
		for col := 0 ;col<lengthCol;col++ {
			if grid[row][col]==1{
				areaIsland := dfs(row,col)
				if areaIsland>maxCount{
					maxCount=areaIsland
				}
			}
		}
	}

	return maxCount
}
