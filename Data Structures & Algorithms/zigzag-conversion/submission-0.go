func convert(s string, numRows int) string {
	// 1 rows -> return s immediately
	if numRows==1{
		return s
	}

	// create an array of string to store the character as the amount of rows
	rows := make([]string,numRows)
	// looking at the pattern
	// 2 rows -> 0 additional inverts
	// 3 rows -> 1 additional inverts
	// 4 rows -> 2 additional inverts
	// 5 rows -> 3 additional inverts
	// using modulo of rows + (rows-2) we can get the amount of going down and up
	
	for i:=0;i<len(s);i++{
		// using modulo of rows + (rows-2) 
		currentGroup := i%(numRows+numRows-2) // 5-> 8 -> 0-7 -> 0-4 
		// if the value <= number of rows -> insert the char at array[modulo%rows]
		if currentGroup < numRows {
			rows[currentGroup%numRows]=rows[currentGroup%numRows]+string(s[i])
		} else {
			// if the value > number of rows -> insert the char at array[rows-modulo%rows]
			// 5 -> 8 -> modulo 5-7 -> 5-7 % 5 -> 0,1,2 -> 5,4,3 -> 3,2,1
			rows[numRows-(currentGroup%numRows)-2] = rows[numRows-(currentGroup%numRows)-2] + string(s[i])
		}
	}
	// loop through array and combine each value
	var result string
	for i:=0;i<numRows;i++{
		result += rows[i]
	}
	// return the value -> O(n)
	return result
}
