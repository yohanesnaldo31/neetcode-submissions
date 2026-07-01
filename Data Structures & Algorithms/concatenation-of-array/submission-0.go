func getConcatenation(nums []int) []int {
    result :=make([]int,2*len(nums))
	for i:=0;i<len(result);i++{
		result[i]=nums[i%len(nums)]
	}
	return result
}
