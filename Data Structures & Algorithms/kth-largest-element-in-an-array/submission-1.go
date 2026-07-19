func findKthLargest(nums []int, k int) int {
	heapArr := []int{0}
	for i:=0;i<len(nums);i++{
		heapArr = append(heapArr,nums[i])
		currPost := len(heapArr)-1 // set current position to the last element index
		for currPost/2 > 0 {
			if heapArr[currPost/2] > heapArr[currPost] {
				heapArr[currPost/2],heapArr[currPost] = heapArr[currPost],heapArr[currPost/2] // swap position
			} else {
				break
			}
			currPost = currPost/2
		}

		if len(heapArr)-1 > k { // if over capaicty, pop the minimum value
			heapArr[1] = heapArr[len(heapArr)-1] 
			heapArr = heapArr[:len(heapArr)-1]

			// check to each children down
			idxStart := 1
			for idxStart*2 < len(heapArr){
				if idxStart*2+1 < len(heapArr) &&  // right child exists
					heapArr[idxStart*2]>heapArr[idxStart*2+1] && // right child value smaller than left child
					heapArr[idxStart*2+1]<heapArr[idxStart]{ // right child value smaller than parrent
					heapArr[idxStart],heapArr[idxStart*2+1] = heapArr[idxStart*2+1],heapArr[idxStart]
					idxStart = idxStart*2+1
				} else if heapArr[idxStart]>heapArr[idxStart*2] {
					heapArr[idxStart],heapArr[idxStart*2] = heapArr[idxStart*2],heapArr[idxStart]
					idxStart = idxStart*2
				} else {
					break
				}
			}
		}
	}
	return heapArr[1]
}
