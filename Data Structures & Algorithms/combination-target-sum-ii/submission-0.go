import (
    "slices"
)

func combinationSum2(candidates []int, target int) [][]int {
    result := make([][]int, 0)
    slices.Sort(candidates)

    var dfsFunc func(currIndex int, currCombination []int, currCombTarget int)
    dfsFunc = func(currIndex int, currCombination []int, currCombTarget int) {
        if currCombTarget == target {
            tempRes := make([]int, len(currCombination))
            copy(tempRes, currCombination)
            result = append(result, tempRes)
            return
        }
        if currCombTarget > target || currIndex >= len(candidates) {
            return
        }

        // Include current element
        currCombination = append(currCombination, candidates[currIndex])
        dfsFunc(currIndex+1, currCombination, currCombTarget+candidates[currIndex])
        currCombination = currCombination[:len(currCombination)-1]

        // Skip duplicates and skip current element
        for currIndex < len(candidates)-1 && candidates[currIndex] == candidates[currIndex+1] {
            currIndex++
        }

        dfsFunc(currIndex+1, currCombination, currCombTarget)
    }

    dfsFunc(0, []int{}, 0)
    return result
}