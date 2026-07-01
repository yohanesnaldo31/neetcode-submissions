func topKFrequent(nums []int, k int) []int {
	result := []int{}
	maps := map[int]int{}
	for _, val := range nums {
		maps[val]++
	}
	arr := make([][]int, len(nums)+1)
	for num, freq := range maps {
		arr[freq] = append(arr[freq], num)
	}
	clear(maps)
	j := len(nums)
	for j >= 0 {
		result = append(result,arr[j]...)
		if len(result) == k {
			break
		}
		j--
	}
	return result
}
