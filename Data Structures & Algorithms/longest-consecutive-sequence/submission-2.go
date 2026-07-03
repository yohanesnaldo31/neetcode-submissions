func longestConsecutive(nums []int) int {
	// map1 -> store unique value -> to check if it exists 
	// map2 -> which we already traversed
	// max count
	// loop through nums
	// X => the start may not be the smallest
	
	if len(nums) == 0 {
		return 0
	}

	
	// map1 -> store unique value -> to check if it exists 
	mapNum := make(map[int]bool)
	for i:=0;i<len(nums);i++{
		mapNum[nums[i]]=true
	}

	// map2 -> store starting point
	mapStartingNum := make(map[int]struct{})
	// loop the nums find the start
	// n-1 -> not found -> start
	for i:=0;i<len(nums);i++{
		_, ok := mapNum[nums[i]-1]
		if !ok {
			mapStartingNum[nums[i]]=struct{}{}
		}
	}

	maxCount := 1
	// iterate from the starting map
	for startingNum, _ := range mapStartingNum{
		count:=1
		currentVal := startingNum
		for mapNum[currentVal+1] == true {
			count++
			currentVal++
		}
		if maxCount<count{
			maxCount=count
		}

	} 

	return maxCount
	// count for each nums + 1 

	// time complexity should be On -> n = number of nums
	// On + Ou + Ou
}
