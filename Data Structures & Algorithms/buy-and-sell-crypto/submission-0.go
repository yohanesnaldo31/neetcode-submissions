func maxProfit(prices []int) int {
	currPoint := 0
	max := 0
	for i:=0;i<len(prices);i++{
		if prices[i]-prices[currPoint] > max {
			max = prices[i]-prices[currPoint]
		}
		if prices[currPoint]>prices[i]{
			currPoint=i
		}
	}
	return max
}
