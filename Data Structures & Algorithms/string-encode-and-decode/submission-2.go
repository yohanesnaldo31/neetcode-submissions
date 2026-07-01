type Solution struct{
	wordLength []int 
}



func (s *Solution) Encode(strs []string) string {
	s.wordLength = make([]int,0)
	encoded := ""
	for _, str := range strs {
		s.wordLength = append(s.wordLength, len(str))
		encoded += str
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string,0)
	currentPointer := 0
	for _, length := range s.wordLength{
		result = append(result, string([]byte(encoded)[currentPointer:currentPointer+length]))
		currentPointer += length
	}
	return result
}
