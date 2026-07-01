func maxArea(heights []int) int {
	idxA := 0;
	idxB := len(heights)-1
	maxArea := 0
	for idxA<idxB {
		area := int(math.Min(float64(heights[idxA]),float64(heights[idxB]))) * (idxB-idxA)
		if maxArea<area{
			maxArea=area
		}

		if heights[idxA]<heights[idxB]{
			idxA++
		} else {
			idxB--
		}
	}
	return maxArea
}
