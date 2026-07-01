func findMaxConsecutiveOnes(nums []int) int {
	var maxConsecutive int
	var currConsecutive int
	for i:=0;i<len(nums);i++{
		if nums[i]==1{
			currConsecutive++
		} else {
			currConsecutive = 0
		}

		if currConsecutive>maxConsecutive{
			maxConsecutive=currConsecutive
		}
	}

	return maxConsecutive
}
