package main

import (
	"errors"
	"fmt"
)

var (
	// ErrEmpty 栈为空
	ErrEmpty = errors.New("stack is empty")
)

// StackNode 栈节点
type StackNode struct {
	el   interface{}
	next *StackNode
}

// LinkStack 栈空间
type LinkStack struct {
	top   *StackNode
	count int
}

// NewChainStack 新建栈
func NewChainStack() *LinkStack {
	return new(LinkStack)
}

// Push 入栈
func (l *LinkStack) Push(el interface{}) (err error) {
	s := StackNode{el: el, next: l.top}
	l.top = &s
	l.count++
	return
}

// Pop 出栈
func (l *LinkStack) Pop() (el interface{}, err error) {
	if l.count == 0 {
		return el, ErrEmpty
	}
	l.count--
	el = l.top.el
	l.top = l.top.next
	return
}

//PrintStack 打印栈中的所有元素
func (l *LinkStack) PrintStack() {
	for i := l.top; i != nil; i = i.next {
		fmt.Print(i.el, " ")
	}
	fmt.Println()
}

// Length 栈长度
func (l *LinkStack) Length() int {
	return l.count
}

// Empty 栈是否为空
func (l *LinkStack) Empty() bool {
	if l.count == 0 {
		return true
	}
	return false
}

// Clear 清空栈
func (l *LinkStack) Clear() {
	l.count = 0
	l.top = nil
}

func main() {
	stack := NewChainStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
}