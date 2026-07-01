func isValid(s string) bool {
	stack := make([]rune,0)
	mapPart := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
    for i:=0;i<len(s);i++{
		if rune(s[i]) == '(' || rune(s[i]) == '[' || rune(s[i]) == '{' {
			stack = append(stack,rune(s[i]))
		} else{
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == mapPart[rune(s[i])] {
				stack=stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack)>0 {
		return false
	}
	return true
}
