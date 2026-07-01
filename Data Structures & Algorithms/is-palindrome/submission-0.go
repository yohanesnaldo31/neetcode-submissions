

func isPalindrome(s string) bool {
    idxA := 0
    idxB := len(s)-1
    str := strings.ToLower(s)
    regexAlphaNumeric := regexp.MustCompile("^[a-zA-Z0-9]+$")
    for idxA<idxB{
        if !regexAlphaNumeric.MatchString(string([]byte(str)[idxA])){
            idxA++
            continue
        }
        if !regexAlphaNumeric.MatchString(string([]byte(str)[idxB])){
            idxB--
            continue
        }
        if []byte(str)[idxA] == []byte(str)[idxB] {
            idxA++
            idxB--
        } else {
            return false
        }
        
    }

    return true
}
