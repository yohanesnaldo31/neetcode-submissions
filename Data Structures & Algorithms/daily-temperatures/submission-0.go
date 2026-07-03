func dailyTemperatures(temperatures []int) []int {
	// loop through temperature -> i
        // loop through temperature + 1 -> j
        // if temperature[i] < temperature[j] 
            // add j-i -> result
            // break
    // On . On-1 -> On2
    // On   


    // ex [3 2 1 1 2 4]
    // for index 0 -> need to know 3
    // the 3 also responsible for index 1 & 2  because of 3
    // need stack
    // 3 2 1 1 -> X -> wat we need is the index
    // 0 1 2 3 -> 4(4)
    // need another stack to store value = stack index -> X we can check the value just by the index
    // else at to stack again  
    // On

    result := make([]int,len(temperatures))
    stack := make([]int,0)
    stack = append(stack,0)
    for i:=1;i<len(temperatures);i++{
        // loop through the stack if index at top value is smaller than current value
        for len(stack)>0 && temperatures[stack[len(stack)-1]]<temperatures[i]{ //
            // val => current index - top index
            result[stack[len(stack)-1]]=i-stack[len(stack)-1]
            // remove every top if current bigger than top
            stack=stack[:len(stack)-1]
        }
        stack=append(stack,i)
    }

    for len(stack)>0{
        result[stack[len(stack)-1]] = 0
        stack=stack[:len(stack)-1]
    }


    return result
}
