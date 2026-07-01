type MinStack struct {
	stack []int
	minVal []int
}

func Constructor() MinStack {
	stack := make([]int,0)
	minVal := make([]int,0)
	return MinStack{
		stack: stack,
		minVal: minVal,
	}
}

func (this *MinStack) Push(val int) {
	this.stack=append(this.stack,val)
	inputtedMinVal := val
	if len(this.minVal)>0{
		currMinVal := this.minVal[len(this.minVal)-1]
		if val>currMinVal{
			inputtedMinVal=currMinVal
		}
	}
	this.minVal=append(this.minVal,inputtedMinVal)
}

func (this *MinStack) Pop() {
	this.stack=this.stack[:len(this.stack)-1]
	this.minVal=this.minVal[:len(this.minVal)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal[len(this.minVal)-1]
}
