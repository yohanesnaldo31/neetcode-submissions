func isAnagram(s string, t string) bool {
	lengthS := len(s)
	lengthT := len(t)
	if lengthS!=lengthT{
		return false
	}
	currentMap := map[byte]int{}
	for i:= 0;i<lengthS;i++{
		currentMap[s[i]]++
	}

	for i:= 0;i<lengthS;i++{
		currentMap[t[i]]--
	}
	for _,val := range currentMap{
		if val !=0 {
			return false
		}
	}
	return true
}
