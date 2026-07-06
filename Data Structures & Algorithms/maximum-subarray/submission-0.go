func maxSubArray(nums []int) int {
    maxSum := nums[0]
	currSum := nums[0]
	for i:=1;i<len(nums);i++{
		if currSum < 0 {
			currSum = 0 
		}

		currSum += nums[i]
		if currSum>maxSum {
			maxSum=currSum
		}
	}
	return maxSum
}
