func threeSum(nums []int) [][]int {
	result := make([][]int,0)
	target := 0
	sort.Slice(nums, func(a,b int)bool{
		return nums[a]<nums[b]
	})

	for i:=0;i<len(nums)-2;i++{
		idxA := i+1
		idxB := len(nums)-1
		for idxA<idxB {
			
			if target-nums[i] > nums[idxA]+nums[idxB] {
				idxA++
				continue
			} else if target-nums[i] < nums[idxA]+nums[idxB] {
				idxB--
				continue
			} else if idxA+1<idxB-1 && nums[idxA]==nums[idxA+1] && nums[idxB]==nums[idxB-1]{ 
				idxA++
				idxB--
			} else {
				result = append(result, []int{nums[i],nums[idxA],nums[idxB]})
				idxA++
				idxB--
			}
		}
		for i<len(nums)-1 && nums[i]==nums[i+1] {
			i++
		}
	}

	return result
}
