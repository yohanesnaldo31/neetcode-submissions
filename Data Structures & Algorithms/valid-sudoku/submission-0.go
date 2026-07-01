func isValidSudoku(board [][]byte) bool {
	mapRows := make([]map[byte]struct{},9) 
	mapColumns := make([]map[byte]struct{},9)
	mapBlocks := make([]map[byte]struct{},9)

	for i:=0;i<9;i++{
		mapRows[i]=make(map[byte]struct{})
		mapColumns[i]=make(map[byte]struct{})
		mapBlocks[i]=make(map[byte]struct{})
	}
	
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			// skip empty
			if board[i][j] == '.' {
				continue
			} 
			// check rows
			_, exists := mapRows[i][board[i][j]]
			if exists {
				return false
			} else {
				mapRows[i][board[i][j]]=struct{}{}
			}

			// check columns
			_, exists = mapColumns[j][board[i][j]]
			if exists {
				return false
			} else {
				mapColumns[j][board[i][j]]=struct{}{}
			}

			// check blocks
			_, exists = mapBlocks[i/3*3+j/3][board[i][j]]
			if exists {
				return false
			} else {
				mapBlocks[i/3*3+j/3][board[i][j]]=struct{}{}
			}
		}
	}

	return true
}
