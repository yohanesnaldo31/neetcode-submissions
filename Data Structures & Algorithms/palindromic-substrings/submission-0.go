func countSubstrings(s string) int {
    if len(s) <= 1 {
        return 1
    }
    result := len(s)

    for i:=1;i<len(s);i++{
        if i+1<len(s) && s[i-1] == s[i+1] {
            idxLeft := i-1
            idxRight := i+1
			result++
            for idxLeft-1>=0 && idxRight+1<len(s) && s[idxLeft-1]==s[idxRight+1]{
                idxLeft--
                idxRight++
				result++
            }
        } 
        if s[i-1]==s[i] {
            idxLeft := i-1
            idxRight := i
			result++
            for idxLeft-1>=0 && idxRight+1<len(s) && s[idxLeft-1] == s[idxRight+1]{
                idxLeft--
                idxRight++
				result++
            }
        } 
    }

    return result
}
