func twoSum(nums []int, target int) []int {
    maps := make(map[int]int)
	for idx, val := range nums {
		_, exists := maps[target-val]
		if !exists {
			maps[val] = idx
		} else {
			return []int{maps[target-val],idx}
		}
	}
	return nil
}
