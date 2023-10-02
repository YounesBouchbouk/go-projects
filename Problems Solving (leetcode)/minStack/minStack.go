package minstack

type MinStack []int

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	*this = append(*this, val)
}

func (this *MinStack) Pop() {
	if len(*this) == 0 {
		return
	}

	indexLastEle := len(*this) - 1
	*this = (*this)[:indexLastEle]
}

func (this *MinStack) Top() int {
	indexLastEle := len(*this) - 1
	ele := (*this)[indexLastEle]
	return ele
}

func (this *MinStack) GetMin() int {
	min := (*this)[0]

	for i := 1; i < len(*this); i++ {
		if (*this)[i] < min {
			min = (*this)[i]
		}
	}
	return min
}
