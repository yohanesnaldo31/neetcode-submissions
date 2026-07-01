func trap(height []int) int {
	idxA := 0
	idxB := len(height)-1
	leftMax := height[0]
	rightMax := height[idxB]
	totalArea := 0
	for idxA < idxB {
		if height[idxA]<height[idxB]{
			idxA++
			if leftMax < height[idxA] {
				leftMax = height[idxA]
			}
			area := int(math.Min(float64(leftMax),float64(rightMax))) - height[idxA]
			if area < 0 {
				area=0
			}
			totalArea+=area
		} else {
			idxB--
			if rightMax < height[idxB] {
				rightMax = height[idxB]
			}
			area := int(math.Min(float64(leftMax),float64(rightMax))) - height[idxB]
			if area < 0 {
				area=0
			}
			totalArea+=area
		}
	}

	return totalArea
}
