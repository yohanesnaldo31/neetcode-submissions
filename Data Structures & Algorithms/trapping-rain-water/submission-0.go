func trap(height []int) int {
	maxRightArr := make([]int,len(height))
	maxLeftArr := make([]int,len(height))

	maxVal := 0
	for idx, num := range height {
		if maxVal <= num {
			maxVal = num
		}
		maxRightArr[idx] = maxVal
	}
	maxVal = 0 
	for i:=len(height)-1;i>=0;i--{
		if maxVal <= height[i] {
			maxVal = height[i]
		}
		maxLeftArr[i] = maxVal
	}

	totalArea := 0
	for i:=0;i<len(height);i++{
		area := int(math.Min(float64(maxRightArr[i]),float64(maxLeftArr[i]))) - height[i]
		if area <0 {
			area = 0
		}
		totalArea+=area
	}

	return totalArea
}
