package main

type CQueue struct {
	in  []int
	out []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 && len(this.in) == 0 {
		return -1
	}

	// 如果只有栈out空,依次弹出 in 栈栈顶元素，入 out 栈
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			lastIndex := len(this.in) - 1
			popVal := this.in[lastIndex]
			this.in = this.in[:lastIndex]
			this.out = append(this.out, popVal)
		}
	}

	// 如果栈out非空
	lastIndex := len(this.out) - 1
	popVal := this.out[lastIndex]
	this.out = this.out[:lastIndex]
	return popVal
}
