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
	for k > 0 && j >= 0 {
		if len(arr[j]) == 0 {
			j--
			continue
		}
		result = append(result,arr[j]...)
		k = k - len(arr[j])
		j--
	}
	return result
}
