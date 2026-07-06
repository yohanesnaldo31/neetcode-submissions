func longestPalindrome(s string) string {
    // init result string
    // loop through
    // find the start of palindrome
    // if n-1==n
    //  idxleft =n-1 , idxRight = n 
    //  for idxLeft-1 == idxRight+1 || still within the range
    //      idxLeft--; idxRight++;
    //  if n-1==n+1
    //  idxleft =n-1 , idxRight = n+1 
    //  for idxLeft-1 == idxRight+1 || still within the range
    //      idxLeft--; idxRight++;
    // check the string index with the result and update result if length > result's
    // return result

    if len(s) <= 1 {
        return s
    }
    result := string(s[0])
    for i:=1;i<len(s);i++{
        if i+1<len(s) && s[i-1] == s[i+1] {
            idxLeft := i-1
            idxRight := i+1
            for idxLeft-1>=0 && idxRight+1<len(s) && s[idxLeft-1]==s[idxRight+1]{
                idxLeft--
                idxRight++
            }
            possibleSubString :=s[idxLeft:idxRight+1]
            if len(result) < len(possibleSubString) {
                result=possibleSubString
            }
        } 
        if s[i-1]==s[i] {
            idxLeft := i-1
            idxRight := i
            for idxLeft-1>=0 && idxRight+1<len(s) && s[idxLeft-1] == s[idxRight+1]{
                idxLeft--
                idxRight++
            }
            possibleSubString :=s[idxLeft:idxRight+1]
            if len(result) < len(possibleSubString) {
                result=possibleSubString
            }
        } 
    }

    return result
}
