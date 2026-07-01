func hasDuplicate(nums []int) bool {
    maps := map[int]struct{}{}
	for _, v:= range nums{
		_, found := maps[v]
		if found {
			return true
		}
		maps[v]= struct{}{}
	}
	return false
}
