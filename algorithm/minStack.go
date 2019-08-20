package main

import "fmt"

/*
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
*/

type MinStack struct {
	elems []int
	mins  []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.elems = append(this.elems, x)

	if len(this.mins) == 0 || this.GetMin() >= x {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	elem := this.Top()
	this.elems = this.elems[:len(this.elems)-1]

	if elem <= this.GetMin() {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.elems) == 0 {
		panic("empty stack")
	}

	elem := this.elems[len(this.elems)-1]
	return elem
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		panic("empty stack")
	}

	elem := this.mins[len(this.mins)-1]
	return elem
}

func main() {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.elems)
	fmt.Println(minStack.mins)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
