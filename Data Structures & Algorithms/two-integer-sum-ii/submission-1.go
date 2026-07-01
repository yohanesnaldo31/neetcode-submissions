func twoSum(numbers []int, target int) []int {
	idxA := 0
	idxB := len(numbers)-1
	for idxA<idxB{
		if numbers[idxA] + numbers[idxB] == target{
			return []int{idxA+1,idxB+1}
		} else if numbers[idxA] + numbers[idxB] > target {
			idxB--
		} else {
			idxA++
		}

	}
	return nil
}
