import (
	"slices"
)

func longestConsecutive(nums []int) int {
	mapNum := make(map[int]int)
	for _,num:=range nums{
		mapNum[num]=1
	}

	arrUnique := make([]int,0)
	for key, _ := range mapNum {
		arrUnique = append(arrUnique,key)
	}

	slices.Sort(arrUnique)
	maxSeq := 0
	for i:=0;i<len(arrUnique);i++{
		_, exists := mapNum[arrUnique[i]+1]
		if exists {
			mapNum[arrUnique[i]+1] = mapNum[arrUnique[i]]+1
		}
	}
	for _, val := range mapNum{
		if val>maxSeq{
			maxSeq=val
		}
	}

	return maxSeq
}
