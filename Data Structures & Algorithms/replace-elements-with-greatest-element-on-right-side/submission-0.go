func replaceElements(arr []int) []int {
	result := make([]int,len(arr))
	biggestVal := -1
	for i:=len(arr)-1;i>=0;i--{
		tempBigVal := arr[i]
		result[i] = biggestVal
		if tempBigVal > biggestVal{
			biggestVal = tempBigVal
		}
	}
	return result
}
