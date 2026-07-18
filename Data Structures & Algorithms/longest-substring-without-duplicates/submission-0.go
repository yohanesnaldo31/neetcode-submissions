func lengthOfLongestSubstring(s string) int {
    if len(s) <= 1 {
        return len(s)
    }

    maxLength := 1
    mapHash := make(map[byte]int)

    leftIndex := 0 
    rightIndex := 1
    mapHash[s[leftIndex]] = leftIndex
    for rightIndex < len(s) {
        existingIdx,  exists := mapHash[s[rightIndex]]
        if !exists {
            currLength := rightIndex-leftIndex + 1
            if maxLength < currLength {
                maxLength = currLength
            }
        } else {
            for leftIndex<=existingIdx{
                delete(mapHash,s[leftIndex])
                leftIndex++
            }
        }
        mapHash[s[rightIndex]] = rightIndex
        rightIndex++
    }
    return maxLength
}