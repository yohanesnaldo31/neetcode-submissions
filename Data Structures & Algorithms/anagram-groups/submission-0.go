import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)
	for _, str := range strs {
		strBytes := []byte(str)
		slices.Sort(strBytes) 
		maps[string(strBytes)] = append(maps[string(strBytes)],str)
	}
	result := [][]string{}
	for _, val := range maps{
		result = append(result,val)
	}
	return result
}
