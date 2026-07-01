func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2{
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	dp := []int{nums[0],nums[1]}
	for i:=2;i<len(nums);i++{
		var opt1 int
		if len(dp)-2>=0{
			opt1 = nums[i]+dp[len(dp)-2]
		}
		var opt2 int
		if len(dp)-3 >= 0{
			opt2 =  nums[i]+dp[len(dp)-3]
		}
		currentVal := max(opt1,opt2)
		dp=append(dp,currentVal)
		if len(dp)>4 {
			dp=dp[1:len(dp)]
		}
	}
	return max(dp[len(dp)-1],dp[len(dp)-2])
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}