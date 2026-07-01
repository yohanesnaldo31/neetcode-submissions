func removeElement(nums []int, val int) int {
    var result,indexTemp int
	tempNums := make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		if nums[i] != val{
			result++
			tempNums[indexTemp]=nums[i]
			indexTemp++
		}
	}
	copy(nums,tempNums)
	return result
}
