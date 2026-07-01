func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	
	dp := []int{1,2}
	
	for i:=3;i<n;i++{
		dp = append(dp, dp[0]+dp[1])
		dp = dp[1:len(dp)]
	}

	return dp[0]+dp[1]
}
