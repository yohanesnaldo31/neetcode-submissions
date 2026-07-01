func largestGoodInteger(num string) string {
	if strings.Contains(num,"999") {
		return "999"
	}
	if strings.Contains(num,"888") {
		return "888"
	}
	if strings.Contains(num,"777") {
		return "777"
	}
	if strings.Contains(num,"666") {
		return "666"
	}
	if strings.Contains(num,"555") {
		return "555"
	}
	if strings.Contains(num,"444") {
		return "444"
	}
	if strings.Contains(num,"333") {
		return "333"
	}
	if strings.Contains(num,"222") {
		return "222"
	}
	if strings.Contains(num,"111") {
		return "111"
	}
	if strings.Contains(num,"000") {
		return "000"
	}
	return ""
}
