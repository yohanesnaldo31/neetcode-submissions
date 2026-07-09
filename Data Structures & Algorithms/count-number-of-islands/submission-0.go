func numIslands(grid [][]byte) int {
	traversedIsland := map[string]struct{}{}
	lengthRow := len(grid)
	lengthCol := len(grid[0])
	var dfs func(indexRow,indexCol int) 
	dfs = func(indexRow,indexCol int) {
		if indexRow<0 || indexCol<0 || indexRow>=lengthRow || indexCol >= lengthCol || grid[indexRow][indexCol]=='0' {
			return
		}
		_, ok := traversedIsland[strconv.Itoa(indexRow)+"#"+strconv.Itoa(indexCol)]
		if ok {
			return
		} 

		traversedIsland[strconv.Itoa(indexRow)+"#"+strconv.Itoa(indexCol)]=struct{}{}
		dfs(indexRow+1,indexCol)
		dfs(indexRow,indexCol+1)
		dfs(indexRow-1,indexCol)
		dfs(indexRow,indexCol-1)
	}
	islandCount := 0
	for row:=0;row<lengthRow;row++{
		for col:=0;col<lengthCol;col++{
			if grid[row][col]=='0'{
				continue
			}
			_, ok := traversedIsland[strconv.Itoa(row)+"#"+strconv.Itoa(col)]
			if ok {
				continue
			} else {
				islandCount++
				dfs(row,col)
			}
		}
	}

	return islandCount
}
