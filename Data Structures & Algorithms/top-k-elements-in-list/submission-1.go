func topKFrequent(nums []int, k int) []int {
	result := []int{}
	maps := map[int]int{}
	for _, val := range nums {
		maps[val]++
	}
	lists := []struct{
		num int
		freq int
	}{}
	for number, occurrence := range maps{
		lists = append(lists,struct{
			num int
			freq int
		}{
			num: number,
			freq: occurrence,
		})
	}
	sort.Slice(lists, func(i, j int) bool {
		return lists[i].freq > lists[j].freq
	})
	for i:=0;i<k;i++{
		result = append(result, lists[i].num)
	}
	return result
}
