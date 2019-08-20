package main

import "fmt"

type StackQueue struct {
	stackA []int
	stackB []int
}

func NewStackQueue() StackQueue {
	return StackQueue{make([]int, 0), make([]int, 0)}
}

//enQueue 入队操作
func (this *StackQueue) enQueue(element int) {
	this.stackA = append(this.stackA, element)
}

//deQueue 出队操作
func (this *StackQueue) deQueue() int {
	element := 0

	if len(this.stackB) == 0 {
		if len(this.stackA) == 0 {
			panic("empty queue")
		}

		//栈A元素转移到栈B
		for len(this.stackA) != 0 {
			this.stackB = append(this.stackB, this.stackA[len(this.stackA) - 1])
			this.stackA = this.stackA[:len(this.stackA) - 1]
		}
	}

	element = this.stackB[len(this.stackB) - 1]
	this.stackB = this.stackB[:len(this.stackB) - 1]

	return element
}

func main() {
	stackQueue := NewStackQueue()
	stackQueue.enQueue(1)
	stackQueue.enQueue(2)
	stackQueue.enQueue(3)
	fmt.Println(stackQueue.stackA, stackQueue.stackB)
	fmt.Println(stackQueue.deQueue())
	fmt.Println(stackQueue.deQueue())
	stackQueue.enQueue(4)
	fmt.Println(stackQueue.stackA, stackQueue.stackB)
	fmt.Println(stackQueue.deQueue())
}