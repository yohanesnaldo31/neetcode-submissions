func orangesRotting(grid [][]int) int {
    // first we need to check if fresh orange exists or not, if not return 0
    // then loop through the grids, find the locations of all rotting oranges and put it in queue oranges
    // we'll store the coordinates in "x#y"
    // we'll use BFS approach starting from all the locations of the first founded rotting oranges
    // we'll try to spread to all of the adjacent oranges 
    // if 0 or 2, then we add nothing to the rotting queue
    // if 1, then we change the value into 2 and put it in to the rotting queue and update the freshOrange counter by -1
    // once all of the first founded rotting oranges executed, check if there is additional into the queue, else we add 1 minute to the result and continue until there is nothing left in the queue
    // if the queue empty and the freshOrange counter is still not 0, then return -1
    // Otherwise, return the minute

    // time Complexity is O(m*n) where m is number of rows and n is number of columns of the grid
    // space complexity is O(m*n) where m is number of rows and n is number of columns of the grid
    
    rottingQueue := make([]string,0)
    freshOrange := 0

    for row:=0;row<len(grid);row++{
        for col:=0;col<len(grid[row]);col++ {
            if grid[row][col] == 1 {
                freshOrange++
            } else if grid[row][col]==2 {
                rottingQueue=append(rottingQueue, strconv.Itoa(row)+"#"+strconv.Itoa(col))
            }
        }
    }

    if freshOrange < 1 {
        return 0
    }

    totalMinutes := -1
    for len(rottingQueue)>0 {
        totalMinutes++
        copyRottingQueue := make([]string,len(rottingQueue))
        copy(copyRottingQueue, rottingQueue)
        for _, coordinate := range copyRottingQueue {
            coordinateSplit := strings.Split(coordinate,"#")
            row, _ := strconv.Atoi(coordinateSplit[0])
            col, _ := strconv.Atoi(coordinateSplit[1])

            if col+1 < len(grid[row]) && grid[row][col+1] == 1 {
                grid[row][col+1] = 2
                rottingQueue = append(rottingQueue, strconv.Itoa(row)+"#"+strconv.Itoa(col+1))
                freshOrange--
            } 
            
            if col-1 >= 0 && grid[row][col-1] == 1 {
                grid[row][col-1] = 2
                rottingQueue = append(rottingQueue, strconv.Itoa(row)+"#"+strconv.Itoa(col-1))
                freshOrange--
            }  
            
            if row+1 < len(grid) && grid[row+1][col] == 1 {
                grid[row+1][col] = 2
                rottingQueue = append(rottingQueue, strconv.Itoa(row+1)+"#"+strconv.Itoa(col))
                freshOrange--
            }  
            
            if row-1 >= 0 && grid[row-1][col] == 1 {
                grid[row-1][col] = 2
                rottingQueue = append(rottingQueue, strconv.Itoa(row-1)+"#"+strconv.Itoa(col))
                freshOrange--
            }
            rottingQueue=rottingQueue[1:]
        }
    } 

    if freshOrange > 0 {
        return -1
    }
    return totalMinutes
}