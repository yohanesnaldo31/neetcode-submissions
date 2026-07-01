func productExceptSelf(nums []int) []int {
	prefixProducts := make([]int,len(nums))
	prefixProducts[0] = 1
	for i:= 1 ; i<len(nums);i++{
		prefixProducts[i] = nums[i-1]*prefixProducts[i-1]
	}  
	suffixProducts := make([]int,len(nums))
	suffixProducts[len(nums)-1] = 1
	for i:= len(nums)-2;i>=0;i--{
		suffixProducts[i] = nums[i+1]*suffixProducts[i+1]
	}
	result := make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		result[i]=prefixProducts[i]*suffixProducts[i]
	}

	return result
}
