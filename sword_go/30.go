package swordgo

type MinStack struct {
	st, mn []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.st = append(this.st, x)
	if len(this.mn) == 0 || x <= this.mn[len(this.mn)-1] {
		this.mn = append(this.mn, x)
	}
}

func (this *MinStack) Pop() {
	v := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	if v == this.mn[len(this.mn)-1] {
		this.mn = this.mn[:len(this.mn)-1]
	}
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) Min() int {
	return this.mn[len(this.mn)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
